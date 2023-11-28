/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

type BaseCell struct {
	Val        []string // 字符串数组， 一个数组代表一行
	style      color.Style
	align      Align
	WordWrap   bool
	rowH, colW int
}

func NewCell(in ...interface{}) Cell {
	if len(in) == 0 {
		return NewEmptyCell()
	}

	cell := &BaseCell{Val: make([]string, 0)}
	for _, val := range in {
		switch val.(type) {
		case string:
			cell.Add(strings.Split(val.(string), "\n")...)
		case fmt.Stringer:
			cell.Add(strings.Split(val.(fmt.Stringer).String(), "\n")...)
		case []string:
			cell.Add(val.([]string)...)
		default:
			cell.Add(fmt.Sprintf("%+v", val))
		}
	}
	return cell
}

func (c *BaseCell) String() (out string) {
	var end = ""
	if c.WordWrap {
		end = "\n"
	}

	lines := c.Lines()
	for idx, line := range lines {
		out += line
		if idx != len(lines)-1 {
			out += end
		}
	}
	return out
}

// Lines 返回这个单元格所有标准行
func (c *BaseCell) Lines() []string {
	var out []string
	for _, val := range c.Val {
		cut := SplitWithRealLength(val, c.colW)
		if len(cut) == 0 {
			continue
		}
		for i := 0; i < len(cut); i++ {
			cut[i] = c.style.Sprint(cut[i])
		}
		out = append(out, cut...)
	}
	for len(out) < c.rowH {
		out = append(out, "")
	}
	return out
}

func (c *BaseCell) Add(in ...string) {
	for _, val := range in {
		val = color.ClearCode(val)
		c.Val = append(c.Val, val)
	}

	if !(c.rowH == 0 && c.colW == 0) {
		c.rowH = len(c.Lines())
		return
	}

	maxL := 0
	for _, val := range c.Val {
		l := RealLength(val)
		if l > maxL {
			maxL = l
		}
	}
	c.colW = maxL
	c.rowH = len(c.Lines())
}

func (c *BaseCell) SetAlign(a Align)                { c.align = a }
func (c *BaseCell) Align() Align                    { return c.align }
func (c *BaseCell) SetColor(style color.Style) Cell { c.style = style; return c }
func (c *BaseCell) Color() color.Style              { return c.style }
func (c *BaseCell) SetWordWrap(b bool) Cell         { c.WordWrap = b; return c }
func (c *BaseCell) SetColWidth(w int) Cell          { c.colW = w; return c }
func (c *BaseCell) SetRowHeight(h int) Cell         { c.rowH = h; return c }
func (c *BaseCell) ColWidth() int                   { return c.colW }
func (c *BaseCell) RowHeight() int                  { return c.rowH }
func (c *BaseCell) IsEmpty() bool                   { return len(c.Val) == 0 && len(c.style) == 0 }

type EmptyCell struct {
	BaseCell
}

func NewEmptyCell() Cell { return &EmptyCell{} }

func (c *EmptyCell) Lines() []string {
	if c.rowH == 0 || c.colW == 0 {
		return []string{""}
	}
	var out []string
	for row := 0; row < c.rowH; row++ {
		out = append(out, strings.Repeat(" ", c.colW))
	}
	return out
}

func NewEmptyCells(col int) Cells {
	var out Cells
	for i := 0; i < col; i++ {
		out = append(out, NewEmptyCell())
	}
	return out
}

func (c Cells2D) String() string {
	out := "[\n"
	for i := 0; i < len(c); i++ {
		out += "\t["
		for j := 0; j < len(c[i]); j++ {
			out += "`" + c[i][j].String() + "`, "
		}
		if len(c[i]) > 0 {
			out = out[:len(out)-2] + "]\n"
			continue
		}
		out += "]\n"
	}
	out += "]"
	return out
}
