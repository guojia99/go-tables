package table

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/guojia99/go-tables/table/utils"
)

type Table struct {
	Headers Row
	Body    Rows
	Footers Row
}

func SimpleTable(in interface{}, align Align) (*Table, error) {
	switch utils.ParsingType(in) {
	case utils.Struct:
		return structTable(in, align)
	case utils.Map:
		return mapTable(in, align)
	case utils.MapSlice:
		return mapSliceTable(in, align)
	case utils.StructSlice:
		return structSliceTable(in, align)
	case utils.Slice:
		return sliceTable(in, align)
	case utils.Slice2D:
		return slice2DTable(in, align)
	}
	return nil, errors.New("the data body required to create a new table frame does not support this type")
}

func mapTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{
		Headers: Row{NewBaseCell(ag, "key"), NewBaseCell(ag, "value")},
	}
	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}
	for _, val := range inValue.MapKeys() {
		tb.Body = append(tb.Body, Row{
			NewBaseCell(ag, fmt.Sprintf("%v", utils.ValueInterface(val))),
			NewBaseCell(ag, fmt.Sprintf("%v", utils.ValueInterface(inValue.MapIndex(val)))),
		})
	}
	return tb, nil
}

func mapSliceTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{}
	inValue := reflect.ValueOf(in)
	keys := inValue.MapKeys()
	maxIdx := 0

	m := make([]reflect.Value, len(keys))
	for idx, key := range keys {
		tb.Headers = append(tb.Headers, NewBaseCell(ag, fmt.Sprintf("%v", utils.ValueInterface(key))))
		v := inValue.MapIndex(key)
		if l := v.Len(); maxIdx < l {
			maxIdx = l
		}
		m[idx] = v
	}
	tb.Body = make(Rows, len(keys))
	for idx := range tb.Body {
		tb.Body[idx] = make(Row, maxIdx)
	}

	for i, val := range m {
		for j := 0; j < maxIdx; j++ {
			if j >= val.Len() {
				tb.Body[i][j] = NewEmptyCell(0, 1)
				continue
			}
			tb.Body[i][j] = NewBaseCell(ag, fmt.Sprintf("%v", utils.ValueInterface(val.Index(j))))
		}
	}
	return tb, nil
}

func structTable(in interface{}, ag Align) (*Table, error) {
	names, value, err := structToRows(in, ag)
	if err != nil {
		return nil, err
	}
	tb := &Table{
		Headers: Row{NewBaseCell(ag, "#"), NewBaseCell(ag, "value")},
	}
	for idx := range names {
		tb.Body = append(tb.Body, Row{names[idx], value[idx]})
	}
	return tb, nil
}

func structSliceTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{}
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}
	for idx, s := range structs {
		names, value, err := structToRows(s, ag)
		if err != nil {
			return nil, err
		}
		if idx == 0 {
			tb.Headers = append(tb.Headers, names...)
		}
		tb.Body = append(tb.Body, value)
	}
	return tb, nil
}

func sliceTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{}
	tb.Headers = append(tb.Headers, NewBaseCell(ag, "No"), NewBaseCell(ag, "value"))
	row, err := sliceToRow(in, ag)
	if err != nil {
		return nil, err
	}
	for idx, val := range row {
		tb.Body = append(tb.Body, Row{
			NewBaseCell(ag, fmt.Sprintf("%d", idx)),
			val,
		})
	}
	return tb, nil
}

func slice2DTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{}
	inValue := reflect.ValueOf(in)
	slice2D := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		slice2D[i] = inValue.Index(i).Interface()
	}

	return tb, nil
}

func structToRows(in interface{}, ag Align) (names, value Row, err error) {
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
		colValue := fmt.Sprintf("%v", utils.ValueInterface(inValue.FieldByName(baseName)))
		if tableTag != "" {
			baseName = tableTag
		} else if jsonTag != "" {
			baseName = jsonTag
		}
		names = append(names, NewBaseCell(ag, baseName))
		value = append(value, NewBaseCell(ag, colValue))
	}
	return
}
func sliceToRow(in interface{}, ag Align) (value Row, err error) {
	inValue := reflect.ValueOf(in)
	row := Row{}
	for i := 0; i < inValue.Len(); i++ {
		val := reflect.ValueOf(inValue.Index(i).Interface())
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		row = append(row, NewBaseCell(ag, fmt.Sprintf("%v", utils.ValueInterface(val))))
	}
	return row, nil
}
