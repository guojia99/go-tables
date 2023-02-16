package table

import (
	`errors`
	`fmt`
	`reflect`
)

type TBKind int

const (
	None          TBKind = iota // other
	IteratorSlice               // Iterator
	CellSlice                   // [] Cell

	String      // string
	Struct      // struct{}
	StructSlice // []struct{}
	Slice       // []interface{}
	Slice2D     // [][]interface{}
	Map         // map[interface{}]interface{}
	MapSlice    // map[interface{}][]interface{}
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
		// todo: 这里应该用取值的方式去拿到一个Cell实际接口才能确认到这个列表是一个[]Cell
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

func parseString(in interface{}) (row []Cell, err error) {
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

func parseStruct(in interface{}) (header, row []Cell, err error) {
	inValue := reflect.ValueOf(in)
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

func parseStructSlice(in interface{}) (header []Cell, body []Cell, err error) { return nil, nil, nil }
func parseSlice(in interface{}) ([]Cell, error)                               { return nil, nil }
func parseSlice2D(in interface{}) ([][]Cell, error)                           { return nil, nil }
func parseMap(in interface{}) (header []Cell, body [][]Cell, err error)       { return nil, nil, nil }
func parseMapSlice(in interface{}) (header []Cell, body [][]Cell, err error)  { return nil, nil, nil }
