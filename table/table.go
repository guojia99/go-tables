package table

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/guojia99/go-tables/table/utils"
)

type Option struct {
	TransformContents TransformContents
	Align             Align
}

type Table struct {
	Headers Row
	Body    Rows
	Footers Row
}

func (o *Option) SimpleTable(in interface{}) (*Table, error) {
	switch utils.ParsingType(in) {
	case utils.Struct:
		return o.structTable(in)
	case utils.Map:
		return o.mapTable(in)
	case utils.MapSlice:
		return o.mapSliceTable(in)
	case utils.StructSlice:
		return o.structSliceTable(in)
	case utils.Slice:
		return o.sliceTable(in)
	case utils.Slice2D:
		return o.slice2DTable(in)
	}
	return nil, errors.New("the data body required to create a new table frame does not support this type")
}

func (o *Option) valueToInterface(in reflect.Value) interface{} {
	return o.TransformContents.Convert(utils.ValueInterface(in))
}

func (o *Option) mapTable(in interface{}) (*Table, error) {
	tb := &Table{
		Headers: Row{NewBaseCell(o.Align, "key"), NewBaseCell(o.Align, "value")},
	}
	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}
	for _, val := range inValue.MapKeys() {
		tb.Body = append(tb.Body, Row{
			NewBaseCell(o.Align, fmt.Sprintf("%v", o.valueToInterface(val))),
			NewBaseCell(o.Align, fmt.Sprintf("%v", o.valueToInterface(inValue.MapIndex(val)))),
		})
	}
	return tb, nil
}

func (o *Option) mapSliceTable(in interface{}) (*Table, error) {
	tb := &Table{}
	inValue := reflect.ValueOf(in)
	keys := inValue.MapKeys()
	maxIdx := 0

	m := make([]reflect.Value, len(keys))
	for idx, key := range keys {
		tb.Headers = append(tb.Headers, NewBaseCell(o.Align, fmt.Sprintf("%v", o.valueToInterface(key))))
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
			tb.Body[i][j] = NewBaseCell(o.Align, fmt.Sprintf("%v", o.valueToInterface(val.Index(j))))
		}
	}
	return tb, nil
}

func (o *Option) structTable(in interface{}) (*Table, error) {
	names, value, err := o.structToRows(in)
	if err != nil {
		return nil, err
	}
	tb := &Table{
		Headers: Row{NewBaseCell(o.Align, "#"), NewBaseCell(o.Align, "value")},
	}
	for idx := range names {
		tb.Body = append(tb.Body, Row{names[idx], value[idx]})
	}
	return tb, nil
}

func (o *Option) structSliceTable(in interface{}) (*Table, error) {
	tb := &Table{}
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}
	for idx, s := range structs {
		names, value, err := o.structToRows(s)
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

func (o *Option) sliceTable(in interface{}) (*Table, error) {
	tb := &Table{}
	tb.Headers = append(tb.Headers, NewBaseCell(o.Align, "No"), NewBaseCell(o.Align, "value"))
	row, err := o.sliceToRow(in)
	if err != nil {
		return nil, err
	}
	for idx, val := range row {
		tb.Body = append(tb.Body, Row{
			NewBaseCell(o.Align, fmt.Sprintf("%d", idx)),
			val,
		})
	}
	return tb, nil
}

func (o *Option) slice2DTable(in interface{}) (*Table, error) {
	tb := &Table{}
	//inValue := reflect.ValueOf(in)
	//slice2D := make([]interface{}, inValue.Len())
	//for i := 0; i < inValue.Len(); i++ {
	//	slice2D[i] = inValue.Index(i).Interface()
	//}
	return tb, nil
}

func (o *Option) structToRows(in interface{}) (names, value Row, err error) {
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
		colValue := fmt.Sprintf("%v", o.valueToInterface(inValue.FieldByName(baseName)))
		if tableTag != "" {
			baseName = tableTag
		} else if jsonTag != "" {
			baseName = jsonTag
		}
		names = append(names, NewBaseCell(o.Align, baseName))
		value = append(value, NewBaseCell(o.Align, colValue))
	}
	return
}

func (o *Option) sliceToRow(in interface{}) (value Row, err error) {
	inValue := reflect.ValueOf(in)
	row := Row{}
	for i := 0; i < inValue.Len(); i++ {
		val := reflect.ValueOf(inValue.Index(i).Interface())
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		row = append(row, NewBaseCell(o.Align, fmt.Sprintf("%v", o.valueToInterface(val))))
	}
	return row, nil
}
