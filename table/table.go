package table

import (
	"github.com/guojia99/go-tables/table/utils"
	"image"
)

type Option struct {
	TransformContents []TransformContent
	Contour           Contour
	Align             Align
}

type tableInterface interface {
	// Copy 拷贝一个新的table，互不干扰
	Copy() *Table
	// String 序列化输出
	String() string
	// AddHeaders 添加头的数据
	AddHeaders(...interface{})
	// AddRow 添加一行的数据
	AddRow(...interface{})
	// AddRowCell 用指定的数据进行添加一行
	AddRowCell(RowCell)
	// GetCellAt 获取某行某列的数据，如果不存在返回 EmptyCell
	GetCellAt(image.Point) Cell
	// SetCellAt 替换某行某列的数据，若超出，则扩容
	SetCellAt(Cell, image.Point)
	// MergeCells 合并某些单元格,注意只会保留最左上角的单元格数据
	MergeCells([]image.Point) error
}

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
