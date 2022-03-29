package table

import (
	"errors"
	"image"

	"github.com/guojia99/go-tables/table/utils"
)

type tableInterface interface {
	// Copy - 拷贝一个新的table，与原有的table互不干扰
	//      - Copy a new table without interfering with the original table
	Copy() *Table
	// String - 序列化输出
	//        - Serialized output
	String() string
	// AddHeaders - 添加头的数据
	//  		  - Add header data
	AddHeaders(...interface{}) *Table
	// SetHeaders - 替换头部
	//            - replace header
	SetHeaders(...interface{}) *Table
	// SetHeadersRow - 替换头部
	//               - RowCell replace header
	SetHeadersRow(RowCell) *Table
	// AddBody - 添加一行的数据
	//    	   - Add a row of data
	AddBody(...interface{}) *Table
	// AddBodyRow - 添加一行的数据
	//    	   - Add a RowCell of data
	AddBodyRow(RowCell) *Table
	// SetBody - 设置某行body的数据
	//  	   - Set the data of a row body
	SetBody(int, ...interface{}) *Table
	// GetCellAt - 获取某行某列的数据，如果不存在返回 EmptyCell
	//  		 - Get the data of a certain row and a certain column, if it does not exist, return EmptyCell
	GetCellAt(image.Point) (Cell, error)
	// SetCellAt - 替换某行某列的数据
	// 			 - Replace the data of a certain row and a certain column
	SetCellAt(image.Point, Cell) error
}

type Option struct {
	Align             Align
	Contour           Contour
	TransformContents []TransformContent
}

type Table struct {
	Opt *Option
	// val
	Headers RowCell
	Body    []RowCell
	Footers RowCell
}

func NewTable(opt *Option) *Table {
	return &Table{
		Opt:     opt,
		Headers: RowCell{},
		Body:    []RowCell{},
		Footers: RowCell{},
	}
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

func (t *Table) String() (out string) {
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
	headerStr := serializedRowCell(tb.Headers, tb.Opt.Contour)
	if headerStr != "" {
		out += headerStr
		out += tb.Opt.Contour.Intersection(maxColWidth)
	}

	for _, val := range tb.Body {
		out += serializedRowCell(val, tb.Opt.Contour)
	}
	out += tb.Opt.Contour.Footer(maxColWidth)
	out += serializedRowCell(tb.Footers, tb.Opt.Contour)
	return
}

func (t *Table) AddHeaders(in ...interface{}) *Table {
	for _, val := range in {
		t.Headers = append(t.Headers, NewInterfaceCell(t.Opt.Align, val))
	}
	return t
}

func (t *Table) SetHeaders(in ...interface{}) *Table {
	var newRows RowCell
	for _, val := range in {
		newRows = append(newRows, NewInterfaceCell(t.Opt.Align, val))
	}
	t.Headers = newRows
	return t
}

func (t *Table) SetHeadersRow(r RowCell) *Table {
	t.Headers = r
	return t
}

func (t *Table) AddBody(in ...interface{}) *Table {
	var newRows RowCell
	for _, val := range in {
		newRows = append(newRows, NewInterfaceCell(t.Opt.Align, val))
	}
	t.Body = append(t.Body, newRows)
	return t
}

func (t *Table) AddBodyRow(r RowCell) *Table {
	t.Body = append(t.Body, r)
	return t
}

func (t *Table) SetBody(idx int, in ...interface{}) *Table {
	if idx >= len(t.Body) {
		return t
	}
	var newRows RowCell
	for _, val := range in {
		newRows = append(newRows, NewInterfaceCell(t.Opt.Align, val))
	}
	t.Body[idx] = newRows
	return t
}

func (t *Table) GetCellAt(p image.Point) (Cell, error) {
	if p.X >= len(t.Body) {
		return nil, errors.New("`row` line out of range")
	}
	if p.Y >= len(t.Body[p.X]) {
		return nil, errors.New("`col` line out of range")
	}
	return t.Body[p.X][p.Y], nil
}

func (t *Table) SetCellAt(p image.Point, c Cell) error {
	switch c.(type) {
	case MergeCell:
		return errors.New("cell can not merge cell")
	}

	if p.X >= len(t.Body) {
		return errors.New("`row` line out of range")
	}
	if p.Y >= len(t.Body[p.X]) {
		return errors.New("`col` line out of range")
	}
	t.Body[p.X][p.Y] = c
	return nil
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

			ws, hs, out = append(ws, addW...), append(hs, addH...), append(out, val)
			continue
		}
		ws, hs, out = append(ws, val.Width()), append(hs, val.Height()), append(out, val)
	}
	return
}
