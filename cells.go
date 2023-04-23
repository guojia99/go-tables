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

type Cells []Cell

func (c Cells) String() string {
	out := ""
	for idx, val := range c {
		out += val.String()
		if idx < len(c)-1 {
			out += ","
		}
	}
	return "[" + out + "]"
}

type Cells2D []Cells

func (c Cells2D) String() string {
	out := ""
	for _, val := range c {
		out += "\t" + val.String()
	}
	return "[\n" + out + "\n]"
}

func NewEmptyCell() Cell { return &BaseCell{} }

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

type Cell interface {
	fmt.Stringer
	Add(...string)
	Lines() []string
	Color() color.Style
	SetColor(color.Style) Cell
	SetWordWrap(b bool) Cell
	SetColWidth(w int)
	SetRowHeight(h int)
	ColWidth() int
	RowHeight() int
}

type BaseCell struct {
	Val        []string
	style      color.Style
	WordWrap   bool
	rowH, colW int
}

func (c *BaseCell) String() (out string) {
	for _, line := range c.Lines() {
		// set colors
		out += c.style.Sprintf("%s", line)
	}
	return out
}
func (c *BaseCell) Lines() []string {
	if c.WordWrap {
		return c.Val
	}
	var out = ""
	for _, val := range c.Val {
		out += c.style.Sprintf("%s", val)
	}
	return []string{out}
}
func (c *BaseCell) Add(in ...string)                { c.Val = append(c.Val, in...) }
func (c *BaseCell) SetColor(style color.Style) Cell { c.style = style; return c }
func (c *BaseCell) Color() color.Style              { return c.style }
func (c *BaseCell) SetWordWrap(b bool) Cell         { c.WordWrap = b; return c }
func (c *BaseCell) SetColWidth(w int)               { c.colW = w }
func (c *BaseCell) SetRowHeight(h int)              { c.rowH = h }
func (c *BaseCell) ColWidth() int                   { return c.colW }
func (c *BaseCell) RowHeight() int                  { return c.rowH }
