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
	- 目前实现了以下几种 Cell, 均基于BaseCell
		= BaseCell
		- 基本单元格
		- 只有最基本的功能

		- EmptyCell
			- 空单元格
			- 没有内容的单元格
		- InterfaceCell
			- 空接口单元格
			- 用于接纳任意类型的单元格
		- MergeCell
			- 合并单元格
			- 可以跨多行多列的单元格
*/
package table

import (
	"fmt"
	"github.com/guojia99/go-tables/table/utils"
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

func (a Align) Repeat(in string, count uint) string {

	if in == "" {
		return strings.Repeat(" ", int(count))
	}

	w := utils.RealLength(in)
	if w < 3 && count > 3 {
		return strings.Repeat(".", int(count))
	}

	repeatLen := int(count) - w

	if repeatLen < 0 {
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

func (a Align) Repeats(in []string, count uint) []string {
	var cache []string
	for _, val := range in {
		if utils.RealLength(val) == 0 {
			continue
		}
		cache = append(cache, a.Repeat(val, count))
	}
	repeatLen := len(in) - len(cache)
	switch a {
	case AlignTopRight, AlignTopLeft, AlignTopCenter:
		return append(cache, make([]string, repeatLen)...)
	case AlignBottomRight, AlignBottomLeft, AlignBottomCenter:
		return append(make([]string, repeatLen), cache...)
	case AlignRight, AlignLeft, AlignCenter:
		leftL := repeatLen / 2
		rightL := repeatLen - leftL
		return append(make([]string, leftL), append(cache, make([]string, rightL)...)...)
	}
	return in
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

func NewBaseCell(ag Align, in []string) *BaseCell {
	c := &BaseCell{
		align: ag,
	}
	c.Add(in...)
	return c
}

func (c *BaseCell) Width() uint      { return c.w }
func (c *BaseCell) Height() uint     { return c.h }
func (c *BaseCell) SetWidth(w uint)  { c.w = w }
func (c *BaseCell) SetHeight(h uint) { c.h = h }
func (c *BaseCell) Align() Align     { return c.align }
func (c *BaseCell) Add(in ...string) {
	for idx, val := range in {
		in[idx] = utils.SwellFont(val)
		if w := uint(utils.RealLength(in[idx])); w > c.w {
			c.w = w
		}
	}
	c.val = append(c.val, in...)
	c.h += uint(len(in))
}
func (c *BaseCell) Line(idx uint) string {
	if idx >= uint(len(c.val)) {
		return c.align.Repeat("", c.w)
	}
	return c.align.Repeat(c.val[idx], c.w)
}
func (c *BaseCell) Lines() (out []string) {
	return c.align.Repeats(c.val, c.w)
}

type EmptyCell struct{ *BaseCell }

func NewEmptyCell(w, h uint) EmptyCell {
	return EmptyCell{
		BaseCell: &BaseCell{
			w: w,
			h: h,
		},
	}
}
func (c EmptyCell) Add(in ...string) { c.h += uint(len(in)) }
func (c EmptyCell) Line(uint) string { return strings.Repeat(" ", int(c.w)) }
func (c EmptyCell) Lines() (out []string) {
	for i := uint(0); i < c.h; i++ {
		out = append(out, c.Line(i))
	}
	return
}

type InterfaceCell struct {
	*BaseCell
	AnyVal []interface{}
}

func (c *InterfaceCell) Add(in ...string) {
	for _, val := range in {
		c.AnyVal = append(c.AnyVal, val)
	}
	c.BaseCell.h = uint(len(c.AnyVal))
}

func (c *InterfaceCell) ToBaseCell() *BaseCell {
	var val []string
	for _, v := range c.AnyVal {
		val = append(val, fmt.Sprintf("%v", v))
	}
	return NewBaseCell(c.align, val)
}

func NewInterfaceCell(ag Align, data ...interface{}) *InterfaceCell {
	val := make([]interface{}, len(data))
	for idx := range data {
		val[idx] = data[idx]
	}

	c := &InterfaceCell{
		BaseCell: &BaseCell{
			align: ag,
		},
		AnyVal: val,
	}
	return c
}

type MergeCell struct {
	*BaseCell
	Row    uint
	Column uint
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
		BaseCell: &BaseCell{
			val:   cells[0][0].Lines(),
			align: cells[0][0].Align(),
		},
	}
}

type RowCell []Cell

func serializedRowCell(r RowCell, c Contour) (out string) {
	if len(r) == 0 {
		return
	}
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
