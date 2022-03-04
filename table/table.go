package table

import (
	"errors"
	"fmt"
	"reflect"
)

type Table struct {
	Headers Row
	Body    Rows
	Footers Row
}

func SimpleTable(in interface{}, align Align) (*Table, error) {
	switch parsingType(in) {
	case Struct:
		return structTable(in, align)
	case Map:
		return mapTable(in, align)
	case MapSlice:
		return mapSliceTable(in, align)
	case StructSlice:
		return structSliceTable(in, align)
	case Slice:
		return sliceTable(in, align)
	case Slice2D:
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
			NewBaseCell(ag, fmt.Sprintf("%v", valueInterface(val))),
			NewBaseCell(ag, fmt.Sprintf("%v", valueInterface(inValue.MapIndex(val)))),
		})
	}
	return tb, nil
}
func mapSliceTable(in interface{}, ag Align) (*Table, error) {
	tb := &Table{}

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
		if !isHeadCapitalLetters(baseName) {
			continue
		}
		tableTag := field.Tag.Get("table")
		jsonTag := field.Tag.Get("json")
		if jsonTag == "-" || tableTag == "-" {
			continue
		}
		colValue := fmt.Sprintf("%v", valueInterface(inValue.FieldByName(baseName)))
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
		row = append(row, NewBaseCell(ag, fmt.Sprintf("%v", valueInterface(val))))
	}
	return row, nil
}
