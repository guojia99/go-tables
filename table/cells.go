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
	Val        []string
	style      color.Style
	WordWrap   bool
	rowH, colW int
}

func NewCell(in ...interface{}) Cell {
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
	// todo 换成 io模型
	for _, line := range c.Lines() {
		out += line
	}
	return out
}
func (c *BaseCell) Lines() []string {
	if c.WordWrap {
		return c.Val
	}
	var out = ""
	for _, val := range c.Val {
		val = color.ClearCode(val)
		out += c.style.Sprintf("%s", val)
	}
	return []string{out}
}
func (c *BaseCell) Add(in ...string)                { c.Val = append(c.Val, in...) }
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
