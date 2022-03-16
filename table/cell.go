/*
	Align 内容方位
	- 以 AlignLeft, AlignCenter, AlignTopLeft 为例,表示在单元格内容分别是置于左侧，和顶部左侧
		= eg. 当单元格高度为4, 实际内容高度为2, 以 AlignLeft 则居左为
			[            ]
			[ align 2    ]
			[ align 3    ]
			[            ]
		= AlignCenter
			[            ]
			[   align 2  ]
			[   align 3  ]
			[            ]
		= AlignTopLeft
			[   align 2  ]
			[   align 3  ]
			[            ]
			[            ]
	- Align.Repeat(in string, w int) string 函数
		= 仅支持针对left, center, rights三个方位的拓展, in 是输入, w是实际所需长度, 若
	Cell 单元格
	- 目前实现了以下几种 Cell
		= BaseCell
			- 基本单元格
			- 只有最基本的功能
		= EmptyCell
			- 空单元格
			- 没有内容的单元格
		= MergeCell
			- 合并单元格
			- 可以跨多行多列的单元格
*/
package table

import (
	"fmt"
	"regexp"
	"runtime"
	"strings"
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight

	AlignTopLeft
	AlignTopCenter
	AlignTopRight

	AlignBottomLeft
	AlignBottomCenter
	AlignBottomRight
)

func (a Align) Repeat(in string, w uint) string {
	if in == "" {
		return strings.Repeat(" ", int(w))
	}

	count := realLength(in)
	if w <= 3 && count >= 3 {
		return strings.Repeat(".", int(w))
	}

	repeatLen := int(w) - count
	if repeatLen <= 0 {
		return in[:int(w)-3] + "..."
	}

	switch a {
	case AlignLeft, AlignTopLeft, AlignBottomLeft:
		return in + strings.Repeat(" ", repeatLen)
	case AlignCenter, AlignTopCenter, AlignBottomCenter:
		leftL := repeatLen / 2
		rightL := repeatLen - leftL
		return strings.Repeat(" ", leftL) + in + strings.Repeat(" ", rightL)
	case AlignRight, AlignTopRight, AlignBottomRight:
		return strings.Repeat(" ", repeatLen) + in
	}
	return strings.Repeat(" ", int(w))
}

type Cell interface {
	Width() uint
	Height() uint
	SetWidth(uint)
	SetHeight(uint)
	Align() Align
	Add(...string)
	Line(uint) string
	Lines() []string
}

type BaseCell struct {
	w, h  uint
	val   []string
	align Align
}

func NewBaseCell(ag Align, data ...string) *BaseCell {
	c := &BaseCell{
		align: ag,
	}
	c.Add(data...)
	return c
}
func (c *BaseCell) Width() uint      { return c.w }
func (c *BaseCell) Height() uint     { return c.h }
func (c *BaseCell) SetWidth(w uint)  { c.w = w }
func (c *BaseCell) SetHeight(h uint) { c.h = h }
func (c *BaseCell) Align() Align     { return c.align }
func (c *BaseCell) Add(in ...string) {
	for idx, val := range in {
		in[idx] = swellFont(val)
		if w := uint(realLength(in[idx])); w > c.w {
			c.w = w
		}
	}
	c.val = append(c.val, in...)
	c.h += uint(len(in))
}
func (c *BaseCell) Line(idx uint) string {
	if idx >= uint(len(c.val)) || idx < 0 {
		return c.align.Repeat("", c.w)
	}
	return c.align.Repeat(c.val[idx], c.w)
}
func (c *BaseCell) Lines() (out []string) {
	for i := uint(0); i < c.h; i++ {
		out = append(out, c.Line(i))
	}
	return
}

type EmptyCell struct{ BaseCell }

func NewEmptyCell(w, h uint) *EmptyCell {
	return &EmptyCell{
		BaseCell{
			w: w,
			h: h,
		},
	}
}
func (c *EmptyCell) Add(in ...string) { c.h += uint(len(in)) }
func (c *EmptyCell) Line(uint) string { return strings.Repeat(" ", int(c.w)) }
func (c *EmptyCell) Lines() (out []string) {
	for i := uint(0); i < c.h; i++ {
		out = append(out, c.Line(i))
	}
	return
}

type MergeCell struct {
	BaseCell
	Row    uint // 行
	Column uint // 列
}

func NewMergeCells(cells [][]Cell) *MergeCell {
	if len(cells) == 0 {
		return &MergeCell{}
	}
	if len(cells[0]) == 0 {
		return &MergeCell{}
	}
	return &MergeCell{
		Row:    uint(len(cells)),
		Column: uint(len(cells[0])),
		BaseCell: BaseCell{
			val:   cells[0][0].Lines(),
			align: cells[0][0].Align(),
		},
	}
}

func swellFont(in string) string { return fmt.Sprintf("%s%s%s", " ", in, " ") }

var noWinCodeExpr = regexp.MustCompile(`\033\[[\d;?]+m`)

func realLength(in string) int {
	if runtime.GOOS != `windows` {
		return stringLength([]rune(noWinCodeExpr.ReplaceAllString(in, "")))
	}
	return stringLength([]rune(in))
}

type FullWidth struct {
	from rune
	to   rune
}

var fullWidth = []FullWidth{
	// Chinese
	{0x2E80, 0x9FD0}, {0xAC00, 0xD7A3}, {0xF900, 0xFACE},
	{0xFE00, 0xFE6C}, {0xFF00, 0xFF60}, {0x20000, 0x2FA1D},
	{12286, 12351},
}

func stringLength(r []rune) (length int) {
	length = len(r)
re:
	for _, val := range r {
		for _, twoBox := range fullWidth {
			if val >= twoBox.from && val <= twoBox.to {
				length++
				continue re
			}
		}
	}
	return
}

type Row []Cell
type Rows []Row

//func MaxH(cells []*Cell) int {
//	mh := 0
//	for _, val := range cells {
//		if val.MaxHeight() > mh {
//			mh = val.MaxHeight()
//		}
//	}
//	return mh
//}
//
//func MaxCellsWidths(cells []*Cell) []int {
//	mws := make([]int, len(cells))
//	for idx, val := range cells {
//		if val.MaxWidth() >= mws[idx] {
//			mws[idx] = val.MaxWidth()
//		}
//	}
//	return mws
//}
//
//type Cells []*Cell
//
//func (cs Cells) Parse(mws []int, opt *Option) string {
//	out := ""
//	if mws == nil || len(mws) != len(cs) {
//		mws = MaxCellsWidths(cs)
//	}
//	for idx, val := range cs {
//		if val.MaxWidth() >= mws[idx] {
//			mws[idx] = val.MaxWidth()
//		}
//	}
//	mh := MaxH(cs)
//	for idx, val := range cs {
//		val.mw = mws[idx]
//		val.mh = mh
//	}
//	for h := 0; h < mh; h++ {
//		out += opt.Contour.V
//		for _, val := range cs {
//			out += val.Line(h) + opt.Contour.V
//		}
//		out += "\n"
//	}
//	return out
//}
