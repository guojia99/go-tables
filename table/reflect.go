/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package table

import (
	`errors`
	`fmt`
	`reflect`
	`sort`
)

type TBKind int

const (
	None          TBKind = iota // other
	IteratorSlice               // Iterator
	CellSlice                   // [] Cell
	String                      // string
	Struct                      // struct{}
	StructSlice                 // []struct{}
	Slice                       // []interface{}
	Slice2D                     // [][]interface{}
	Map                         // map[interface{}]interface{}
	MapSlice                    // map[interface{}][]interface{}
)

func (t TBKind) String() string {
	switch t {
	case IteratorSlice:
		return "iterator"
	case Struct:
		return "struct"
	case Map:
		return "map"
	case CellSlice:
		return "cell slice"
	case MapSlice:
		return "map slice"
	case StructSlice:
		return "struct slice"
	case Slice:
		return "slice"
	case Slice2D:
		return "slice 2D"
	}
	return "unknown"
}

// parsingTypeTBKind judgment on whether the parsed data is legal and can be converted into tabular data
func parsingTypeTBKind(in interface{}) TBKind {
	val := reflect.ValueOf(in)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	// check types
	switch val.Interface().(type) {
	case Iterator:
		return IteratorSlice
	case fmt.Stringer:
		return String
	}

	// check kinds is base value kind
	switch val.Kind() {
	case reflect.String:
		return String
	case reflect.Struct:
		return Struct
	case reflect.Map:
		// map type.kind.elem is values -> map[Key]Value
		switch val.Type().Elem().Kind() {
		case reflect.Slice, reflect.Array:
			return MapSlice
		}
		return Map
	case reflect.Slice, reflect.Array:
		// todo: list has use value check
		// todo: 这里应该用取值的方式去拿到一个Cell实际接口才能确认到这个列表是一个Cells
		if val.Type().Elem().String() == "table.Cell" {
			return CellSlice
		}
		// slice type.kind.elem is values -> []Value
		switch val.Type().Elem().Kind() {
		case reflect.Struct:
			return StructSlice
		case reflect.Slice, reflect.Array:
			return Slice2D
		}
		return Slice
	}
	return None
}

func parseString(in interface{}) (row Cells, err error) {
	switch in.(type) {
	case string:
		row = append(row, NewCell(in))
	case fmt.Stringer:
		row = append(row, NewCell(in.(fmt.Stringer).String()))
	default:
		err = fmt.Errorf("the data is not a string type data")
	}
	return
}

// parseStruct
// parse the structure, and use the key value content of the structure as the header and row as the return value.
func parseStruct(in interface{}) (header, row Cells, err error) {
	inValue := reflect.ValueOf(in)
	if !inValue.IsValid() {
		err = errors.New("the content of the struct is not valid")
		return
	}

	inType := inValue.Type()
	if inValue.Kind() == reflect.Ptr {
		inValue, inType = inValue.Elem(), inType.Elem()
	}
	if inValue.Kind() != reflect.Struct {
		err = errors.New("the content of the struct list is not a struct")
		return
	}

	// find all has `json`、`table`、`yaml` tag filed to table cells
	for n := 0; n < inValue.NumField(); n++ {
		if filed := inType.Field(n); isHeadCapitalLetters(filed.Name) {
			if baseName, ok := structTagName(filed.Tag); ok {
				header = append(header, NewCell(baseName))
				row = append(row, NewCell(valueInterface(inValue.FieldByName(filed.Name))))
			}
		}
	}
	return
}

func parseStructSlice(in interface{}) (header Cells, body []Cells, err error) {
	inValue := reflect.ValueOf(in)
	if !inValue.IsValid() {
		err = errors.New("the content of the struct slice is not valid")
		return
	}
	if inValue.Kind() != reflect.Slice && inValue.Kind() != reflect.Array {
		err = errors.New("the content is not a struct slice")
		return
	}

	for i := 0; i < inValue.Len(); i++ {
		ptr := inValue.Index(i)
		if inValue.Index(i).Kind() == reflect.Ptr {
			ptr = ptr.Elem()
		}

		h, b, parseStructErr := parseStruct(ptr.Interface())
		if parseStructErr != nil {
			err = fmt.Errorf("parse %d interface error: %s", i, parseStructErr)
			return
		}
		header = h
		body = append(body, b)
	}
	return
}

func parseSlice(in interface{}) (body Cells, err error) {
	inValue := reflect.ValueOf(in)
	if !inValue.IsValid() {
		err = errors.New("the content of the slice is not valid")
		return
	}
	if inValue.Kind() != reflect.Slice && inValue.Kind() != reflect.Array {
		err = errors.New("the content is not a slice")
		return
	}

	for i := 0; i < inValue.Len(); i++ {
		val := reflect.ValueOf(inValue.Index(i).Interface())
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		body = append(body, NewCell(valueInterface(val)))
	}
	return
}

func parseSlice2D(in interface{}) (body []Cells, err error) {
	inValue := reflect.ValueOf(in)
	if !inValue.IsValid() {
		err = errors.New("the content of the slice 2D is not valid")
		return
	}
	if inValue.Kind() != reflect.Slice && inValue.Kind() != reflect.Array {
		err = errors.New("the content is not a slice 2D")
		return
	}

	for i := 0; i < inValue.Len(); i++ {
		slice, parseSliceErr := parseSlice(inValue.Index(i).Interface())
		if err != nil {
			err = fmt.Errorf("parse %d interface error: %s", i, parseSliceErr)
			return
		}
		body = append(body, slice)
	}
	return
}

func parseMap(in interface{}) (header Cells, body Cells, err error) {
	inValue := reflect.ValueOf(in)
	if !inValue.IsValid() {
		err = errors.New("the content of the map is not valid")
		return
	}
	if inValue.Kind() != reflect.Map {
		err = errors.New("the content is not a map")
		return
	}

	keys := inValue.MapKeys()
	var maps = make(sortMapKeyValues, len(keys))
	for _, val := range inValue.MapKeys() {
		maps = append(maps, sortMapKeyValue{
			key:   NewCell(valueInterface(val)),
			value: NewCell(valueInterface(inValue.MapIndex(val))),
		})
	}
	sort.Sort(maps)
	for _, val := range maps {
		header = append(header, val.key)
		body = append(body, val.value)
	}
	return
}

func parseMapSlice(in interface{}) (header Cells, body []Cells, err error) {
	return nil, nil, nil
}
