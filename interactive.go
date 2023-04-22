/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
)

func AnyTables(in interface{}) (Table, error) {
	kind := parsingTypeTBKind(in)
	if kind == None {
		return nil, fmt.Errorf("the input data is none table")
	}

	tb := &table{
		headers: make([]Cells, 0),
		footers: make([]Cells, 0),
		body:    make([]Cells, 0),
	}

	switch kind {
	case IteratorSlice:
		tb.iterator = in.(Iterator)
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
		tb.headers = append(tb.headers, header)
		tb.body = append(tb.body, row)
	case StructSlice:
		header, body, err := parseStructSlice(in)
		if err != nil {
			return nil, err
		}
		tb.headers = append(tb.headers, header)
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
		header, row, err := parseMap(in)
		if err != nil {
			return nil, err
		}
		tb.headers = append(tb.headers, header)
		tb.body = append(tb.body, row)
	case MapSlice:
		header, body, err := parseMapSlice(in)
		if err != nil {
			return nil, err
		}
		tb.headers = append(tb.headers, header)
		tb.body = append(tb.body, body...)
	}
	return tb, nil
}
