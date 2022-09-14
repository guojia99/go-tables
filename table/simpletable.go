package table

import (
	"errors"
	"fmt"
	"log"
	"math"
	"reflect"

	"github.com/guojia99/go-tables/table/utils"
)

func DefaultSimpleTable(in interface{}) *Table {
	tb, err := SimpleTable(in, DefaultOption)
	if err != nil {
		log.Default().Println(err)
		return nil
	}
	return tb
}

func SimpleTable(in interface{}, opt *Option) (*Table, error) {
	switch utils.ParsingType(in) {
	case utils.Struct:
		return structTable(in, opt)
	case utils.Map:
		return mapTable(in, opt)
	case utils.MapSlice:
		return mapSliceTable(in, opt)
	case utils.StructSlice:
		return structSliceTable(in, opt)
	case utils.Slice:
		return sliceTable(in, opt)
	case utils.Slice2D:
		return slice2DTable(in, opt)
	}
	return nil, errors.New("the data body required to create a new table frame does not support this type")
}

func mapTable(in interface{}, opt *Option) (*Table, error) {
	tb := NewTable(opt).SetHeaders("key", "value")

	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}
	for _, val := range inValue.MapKeys() {
		tb.AddBody(utils.ValueInterface(val), utils.ValueInterface(inValue.MapIndex(val)))
	}
	return tb, nil
}

func mapSliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := NewTable(opt)
	inValue := reflect.ValueOf(in)
	keys := inValue.MapKeys()

	var (
		maxIdx   float64 = 0
		m                = make([]reflect.Value, len(keys))
		keyNames         = make([]string, len(keys))
	)

	for idx, key := range keys {
		keyInterface := utils.ValueInterface(key)
		tb.AddHeaders(keyInterface)
		KeyName := fmt.Sprintf("%s", keyInterface)
		keyNames = append(keyNames, KeyName)

		slice := inValue.MapIndex(key)
		m[idx] = slice
		maxIdx = math.Max(maxIdx, float64(slice.Len()))
	}

	tb.Body = make([]RowCell, int(maxIdx))
	for idx := range tb.Body {
		tb.Body[idx] = make(RowCell, len(keys))
	}

	for i, val := range m {
		for j := 0; j < int(maxIdx); j++ {
			if j >= val.Len() {
				tb.Body[j][i] = NewEmptyCell(0, 1)
				continue
			}
			tb.Body[j][i] = NewInterfaceCell(opt.Align, utils.ValueInterface(val.Index(j)))
		}
	}

	return tb, nil
}

func structTable(in interface{}, opt *Option) (*Table, error) {
	names, value, err := structToRows(in, opt.Align)
	if err != nil {
		return nil, err
	}
	tb := NewTable(opt).SetHeaders("key", "value")
	for idx := range names {
		tb.AddBodyRow(RowCell{names[idx], value[idx]})
	}
	return tb, nil
}

func structSliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := NewTable(opt)
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		data := inValue.Index(i)
		if inValue.Index(i).Kind() == reflect.Ptr {
			structs[i] = data.Elem().Interface()
			continue
		}
		structs[i] = data.Interface()
	}
	for idx, s := range structs {
		names, value, err := structToRows(s, opt.Align)
		if err != nil {
			return &Table{}, err
		}
		if idx == 0 {
			tb.SetHeadersRow(names)
		}
		tb.AddBodyRow(value)
	}
	return tb, nil
}

func sliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := NewTable(opt).AddHeaders("value")
	row, err := sliceToRow(in, opt.Align)
	if err != nil {
		return &Table{}, err
	}
	for _, val := range row {
		tb.AddBody(val)
	}
	return tb, nil
}

func slice2DTable(in interface{}, opt *Option) (*Table, error) {
	tb := NewTable(opt)
	inValue := reflect.ValueOf(in)
	tb.Body = make([]RowCell, inValue.Len())

	var err error
	for i := 0; i < inValue.Len(); i++ {
		slice := inValue.Index(i).Interface()
		tb.Body[i], err = sliceToRow(slice, opt.Align)
		if err != nil {
			return &Table{}, err
		}
	}
	return tb, nil
}

func structToRows(in interface{}, ag Align) (names, value RowCell, err error) {
	inValue := reflect.ValueOf(in)
	if inValue.Kind() != reflect.Struct {
		return nil, nil, errors.New("the content of the struct list is not a struct")
	}

	inType := reflect.TypeOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
		inType = inType.Elem()
	}
	for n := 0; n < inValue.NumField(); n++ {
		field := inType.Field(n)
		baseName := field.Name
		if !utils.IsHeadCapitalLetters(baseName) {
			continue
		}
		tableTag := field.Tag.Get("table")
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" || tableTag == "-" {
			continue
		}
		colValue := inValue.FieldByName(baseName)
		if tableTag != "" {
			baseName = tableTag
		} else if jsonTag != "" {
			baseName = jsonTag
		}
		names = append(names, NewInterfaceCell(ag, baseName))
		value = append(value, NewInterfaceCell(ag, utils.ValueInterface(colValue)))
	}
	return
}

func sliceToRow(in interface{}, ag Align) (value RowCell, err error) {
	inValue := reflect.ValueOf(in)
	row := RowCell{}
	for i := 0; i < inValue.Len(); i++ {
		val := reflect.ValueOf(inValue.Index(i).Interface())
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		row = append(row, NewInterfaceCell(ag, utils.ValueInterface(val)))
	}
	return row, nil
}
