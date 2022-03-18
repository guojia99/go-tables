package table

import (
	"errors"
	"github.com/guojia99/go-tables/table/utils"
	"reflect"
)

type Option struct {
	TransformContents []TransformContent
	Contour           Contour
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

func (t *Table) Copy() *Table {
	newT := &Table{
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

func (t Table) String() (out string) {
	// Make a copy to avoid data confusion
	tb := t.Copy()

	// Get header\footer width parameter
	var headerWidth, headerHeight []uint
	tb.Headers, headerWidth, headerHeight = tb.coverCell(tb.Headers)
	var footerWidth, footerHeight []uint
	tb.Footers, footerWidth, footerHeight = tb.coverCell(tb.Footers)

	// Get body width parameter
	var (
		bodyWidths  = make([][]uint, len(tb.Body))
		bodyHeights = make([][]uint, len(tb.Body))
	)
	for idx := range tb.Body {
		tb.Body[idx], bodyWidths[idx], bodyHeights[idx] = tb.coverCell(tb.Body[idx])
	}

	// Calculate Equilibrium Column Parameters
	maxCol := len(headerWidth)
	for _, bw := range bodyWidths {
		maxCol = utils.Max(maxCol, len(bw))
	}
	var maxColWidth = make([]uint, maxCol)
	for i := 0; i < maxCol; i++ {
		if len(headerWidth) > i && maxColWidth[i] < headerWidth[i] {
			maxColWidth[i] = headerWidth[i]
		}
		if len(footerWidth) > i && maxColWidth[i] < footerWidth[i] {
			maxColWidth[i] = footerWidth[i]
		}
		for bIdx := range bodyWidths {
			if len(bodyWidths[bIdx]) > i && maxColWidth[i] < bodyWidths[bIdx][i] {
				maxColWidth[i] = bodyWidths[bIdx][i]
			}
		}
	}

	// Modify Cell parameters
	for idx := range tb.Headers {
		var saveWidth = maxColWidth[idx]
		switch tb.Headers[idx].(type) {
		case *MergeCell:
			saveWidth = 0
			vM := tb.Headers[idx].(*MergeCell)
			for i := 0; i < int(vM.Column); i++ {
				saveWidth += maxColWidth[i+idx]
			}
		}
		tb.Headers[idx].SetWidth(saveWidth)
		tb.Headers[idx].SetHeight(utils.UintMax(headerHeight...))
	}
	for idx := range tb.Footers {
		var saveWidth = maxColWidth[idx]
		switch tb.Footers[idx].(type) {
		case *MergeCell:
			saveWidth = 0
			vM := tb.Footers[idx].(*MergeCell)
			for i := 0; i < int(vM.Column); i++ {
				saveWidth += maxColWidth[i+idx]
			}
		}
		tb.Footers[idx].SetWidth(saveWidth)
		tb.Footers[idx].SetHeight(utils.UintMax(footerHeight...))
	}

	// Reset line by line
	for rowIdx := range tb.Body {
		for colIdx := range tb.Body[rowIdx] {
			var saveWidth = maxColWidth[colIdx]
			switch tb.Body[rowIdx][colIdx].(type) {
			case *MergeCell:
				saveWidth = 0
				vM := tb.Body[rowIdx][colIdx].(*MergeCell)
				for i := 0; i < int(vM.Column); i++ {
					saveWidth += maxColWidth[i+colIdx]
				}
			}
			tb.Body[rowIdx][colIdx].SetWidth(saveWidth)
			tb.Body[rowIdx][colIdx].SetHeight(utils.UintMax(bodyHeights[rowIdx]...))
		}
	}

	// Serialized output
	out += tb.Opt.Contour.Handler(maxColWidth)
	out += serializedRowCell(tb.Headers, tb.Opt.Contour)
	out += tb.Opt.Contour.Intersection(maxColWidth)
	for _, val := range tb.Body {
		out += serializedRowCell(val, tb.Opt.Contour)
	}
	out += tb.Opt.Contour.Footer(maxColWidth)
	out += serializedRowCell(tb.Footers, tb.Opt.Contour)
	return
}

func serializedRowCell(r RowCell, c Contour) (out string) {
	var heights []uint
	var data [][]string
	for _, val := range r {
		heights = append(heights, val.Height())
		data = append(data, val.Lines())
	}
	maxHeight := utils.UintMax(heights...)
	for idx := 0; idx < int(maxHeight); idx++ {
		out += c.L
		for valIdx, val := range data {
			out += val[idx]
			if valIdx < len(data)-1 {
				out += c.CH
				continue
			}
			out += c.R + "\n"
		}
	}
	return
}

func (t *Table) transformCover(in interface{}) interface{} {
	if len(t.Opt.TransformContents) == 0 {
		return in
	}
	for _, f := range t.Opt.TransformContents {
		in = f(in)
	}
	return in
}

func (t *Table) coverCell(in RowCell) (out RowCell, ws, hs []uint) {
	for _, val := range in {
		switch val.(type) {
		case *InterfaceCell:
			vI := val.(*InterfaceCell)
			var newAnyVal []interface{}
			for i := range vI.AnyVal {
				data := vI.AnyVal[i]
				for _, f := range t.Opt.TransformContents {
					data = f(data)
				}
				newAnyVal = append(newAnyVal, data)
			}
			val = vI.ToBaseCell()
		case *MergeCell:
			vM := val.(*MergeCell)
			addW := make([]uint, vM.Column)
			wc := vM.Width() % vM.Column
			for idx := range addW {
				addW[idx] = vM.Width() / vM.Column
				if uint(idx) < wc {
					addW[idx]++
				}
			}

			addH := make([]uint, vM.Row)
			hr := vM.Height() % vM.Row
			for idx := range addH {
				addH[idx] = vM.Height() / vM.Row
				if uint(idx) < hr {
					addH[idx]++
				}
			}

			ws = append(ws, addW...)
			hs = append(hs, addH...)
			out = append(out, val)
			continue
		}
		ws = append(ws, val.Width())
		hs = append(hs, val.Height())
		out = append(out, val)
	}
	return
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
	return &Table{}, errors.New("the data body required to create a new table frame does not support this type")
}

func mapTable(in interface{}, opt *Option) (*Table, error) {
	tb := &Table{
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

func mapSliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := &Table{Opt: opt}
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

func structTable(in interface{}, opt *Option) (*Table, error) {
	names, value, err := structToRows(in, opt.Align)
	if err != nil {
		return &Table{}, err
	}
	tb := &Table{
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

func structSliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := &Table{Opt: opt}
	inValue := reflect.ValueOf(in)
	structs := make([]interface{}, inValue.Len())
	for i := 0; i < inValue.Len(); i++ {
		structs[i] = inValue.Index(i).Interface()
	}
	for idx, s := range structs {
		names, value, err := structToRows(s, opt.Align)
		if err != nil {
			return &Table{}, err
		}
		if idx == 0 {
			tb.Headers = append(tb.Headers, names...)
		}
		tb.Body = append(tb.Body, value)
	}
	return tb, nil
}

func sliceTable(in interface{}, opt *Option) (*Table, error) {
	tb := &Table{Opt: opt}
	tb.Headers = append(tb.Headers, NewInterfaceCell(opt.Align, "No"), NewInterfaceCell(opt.Align, "value"))
	row, err := sliceToRow(in, opt.Align)
	if err != nil {
		return &Table{}, err
	}
	for idx, val := range row {
		tb.Body = append(tb.Body, RowCell{
			NewInterfaceCell(opt.Align, idx),
			val,
		})
	}
	return tb, nil
}

func slice2DTable(in interface{}, opt *Option) (*Table, error) {
	tb := &Table{Opt: opt}
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
