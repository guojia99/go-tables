package tables

import "fmt"

func anyTables(in interface{}) (Table, error) {
	kind := parsingTypeTBKind(in)
	if kind == None {
		return nil, fmt.Errorf("the input data is none table")
	}

	tb := &table{
		body: make([]Cells, 0),
	}

	switch kind {
	case CellSlice:
		tb.body = append(tb.body, in.([]Cell))
	case String:
		row, err := parseString(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, row)
	case Struct:
		header, row, err := parseStruct(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, header)
		tb.body = append(tb.body, row)
	case StructSlice:
		header, body, err := parseStructSlice(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, header)
		tb.body = append(tb.body, body...)
	case Slice:
		row, err := parseSlice(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, row)
	case Slice2D:
		body, err := parseSlice2D(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, body...)
	case Map:
		header, body, err := parseMap(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, header)
		tb.body = append(tb.body, body)
	case MapSlice:
		header, body, err := parseMapSlice(in)
		if err != nil {
			return nil, err
		}
		tb.body = append(tb.body, header)
		tb.body = append(tb.body, body...)
	}
	return tb, nil
}
