package table

import (
	"errors"
	"fmt"
	"reflect"
	"sort"

	"github.com/guojia99/go-tables/table/utils"
)

type Option struct {
	TransformContents []TransformContent
	Align             Align
}

type RowCell []Cell

type Table struct {
	Opt *Option
	// val
	Headers RowCell
	Body    []RowCell
	Footers RowCell
}

func (t *Table) Copy() Table {
	newT := Table{
		Opt:     t.Opt,
		Body:    make([]RowCell, len(t.Body)),
		Headers: make(RowCell, len(t.Headers)),
		Footers: make(RowCell, len(t.Footers)),
	}
	copy(newT.Body, t.Body)
	copy(newT.Headers, t.Headers)
	copy(newT.Footers, t.Footers)
	return newT
}

func (t *Table) TransformCover(in interface{}) interface{} {
	if len(t.Opt.TransformContents) == 0 {
		return in
	}
	for _, f := range t.Opt.TransformContents {
		in = f(in)
	}
	return in
}

func (t Table) coverCell(in RowCell) (RowCell, []uint) {
	ws := make([]uint, len(in))
	var out RowCell
	for idx, val := range in {
		switch val.(type) {
		case InterfaceCell:
			vI := val.(InterfaceCell)
			var newAnyVal []interface{}
			for i := range vI.AnyVal {
				data := vI.AnyVal[i]
				for _, f := range t.Opt.TransformContents {
					data = f(data)
				}
				newAnyVal = append(newAnyVal, data)
			}
			val = vI.ToBaseCell()
		}
		ws[idx] = val.Width()
		out = append(out, val)
	}
	return out, ws
}

func (t *Table) String() string {
	tb := t.Copy()
	//var (
	//	headerWidth []uint
	//	footerWidth []uint
	//)
	//t.Headers, headerWidth = t.coverCell(t.Headers)
	//t.Footers, footerWidth = t.coverCell(t.Footers)
	//fmt.Println(headerWidth, footerWidth)
	//fmt.Println(t.Headers, t.Footers)
	for idx, bodyCell := range tb.Body {
		bodyCell, bodyCellW := tb.coverCell(bodyCell)
		sort.Slice(bodyCellW, func(i, j int) bool {
			return bodyCellW[i] > bodyCellW[j]
		})
		if len(bodyCellW) == 0 {
			continue
		}
		for _, val := range bodyCell {
			val.SetWidth(bodyCellW[0])
		}
		tb.Body[idx] = bodyCell
	}

	for _, val := range tb.Body {
		for _, v := range val {
			fmt.Println(v.Lines())
		}
	}

	return ""
}

func SimpleTable(in interface{}, opt *Option) (Table, error) {
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
	return Table{}, errors.New("the data body required to create a new table frame does not support this type")
}

func mapTable(in interface{}, opt *Option) (Table, error) {
	tb := Table{
		Opt: opt,
		Headers: RowCell{
			NewInterfaceCell(opt.Align, "key"),
			NewInterfaceCell(opt.Align, "value"),
		},
	}
	inValue := reflect.ValueOf(in)
	if inValue.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
	}
	for _, val := range inValue.MapKeys() {
		tb.Body = append(tb.Body, RowCell{
			NewInterfaceCell(opt.Align, utils.ValueInterface(val)),
			NewInterfaceCell(opt.Align, utils.ValueInterface(inValue.MapIndex(val))),
		})
	}
	return tb, nil
}

func mapSliceTable(in interface{}, opt *Option) (Table, error) {
	tb := Table{Opt: opt}
	inValue := reflect.ValueOf(in)
	keys := inValue.MapKeys()
	maxIdx := 0

	m := make([]reflect.Value, len(keys))
	for idx, key := range keys {
		tb.Headers = append(tb.Headers, NewInterfaceCell(opt.Align, utils.ValueInterface(key)))
		v := inValue.MapIndex(key)
		if l := v.Len(); maxIdx < l {
			maxIdx = l
		}
		m[idx] = v
	}
	tb.Body = make([]RowCell, len(keys))
	for idx := range tb.Body {
		tb.Body[idx] = make(RowCell, maxIdx)
	}

	for i, val := range m {
		for j := 0; j < maxIdx; j++ {
			if j >= val.Len() {
				tb.Body[i][j] = NewEmptyCell(0, 1)
				continue
			}
			tb.Body[i][j] = NewInterfaceCell(opt.Align, utils.ValueInterface(val.Index(j)))
		}
	}
	return tb, nil
}

func structTable(in interface{}, opt *Option) (Table, error) {
	names, value, err := structToRows(in, opt.Align)
	if err != nil {
		return Table{}, err
	}
	tb := Table{
		Opt: opt,
		Headers: RowCell{
			NewInterfaceCell(opt.Align, "#"),
			NewInterfaceCell(opt.Align, "value"),
		},
	}
	for idx := range names {
		tb.Body = append(tb.Body, RowCell{names[idx], value[idx]})
	}
	return tb, nil
}

func structSliceTable(in interface{}, opt *Option) (Table, error) {
	tb := Table{Opt: opt}
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}
	for idx, s := range structs {
		names, value, err := structToRows(s, opt.Align)
		if err != nil {
			return Table{}, err
		}
		if idx == 0 {
			tb.Headers = append(tb.Headers, names...)
		}
		tb.Body = append(tb.Body, value)
	}
	return tb, nil
}

func sliceTable(in interface{}, opt *Option) (Table, error) {
	tb := Table{Opt: opt}
	tb.Headers = append(tb.Headers, NewInterfaceCell(opt.Align, "No"), NewInterfaceCell(opt.Align, "value"))
	row, err := sliceToRow(in, opt.Align)
	if err != nil {
		return Table{}, err
	}
	for idx, val := range row {
		tb.Body = append(tb.Body, RowCell{
			NewInterfaceCell(opt.Align, idx),
			val,
		})
	}
	return tb, nil
}

func slice2DTable(in interface{}, opt *Option) (Table, error) {
	tb := Table{Opt: opt}
	inValue := reflect.ValueOf(in)
	tb.Body = make([]RowCell, inValue.Len())

	var err error
	for i := 0; i < inValue.Len(); i++ {
		slice := inValue.Index(i).Interface()
		tb.Body[i], err = sliceToRow(slice, opt.Align)
		if err != nil {
			return Table{}, err
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
