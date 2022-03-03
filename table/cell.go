package table

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
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

func (a Align) Repeat(in string, w int) string {
	l := w - realLength(in)
	if in == "" || l < 0 {
		return strings.Repeat(" ", w)
	}

	switch a {
	case AlignLeft, AlignTopLeft, AlignBottomLeft:
		return in + strings.Repeat(" ", l)
	case AlignCenter, AlignTopCenter, AlignBottomCenter:
		leftL := l / 2
		rightL := l - leftL
		return strings.Repeat(" ", leftL) + in + strings.Repeat(" ", rightL)
	case AlignRight, AlignTopRight, AlignBottomRight:
		return strings.Repeat(" ", l) + in
	}
	return strings.Repeat(" ", w)
}

func swellFont(in string) string { return fmt.Sprintf("%s%s%s", " ", in, " ") }
func realLength(in string) int   { return stringLength([]rune(color.ClearCode(in))) }

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

type Cell interface {
	Width() int
	Height() int
	SetWidth(int)
	SetHeight(int)
	Add(...string)
	Line(int) string
	Lines() []string
}

type BaseCell struct {
	w, h  int
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
func (c *BaseCell) Width() int      { return c.w }
func (c *BaseCell) Height() int     { return c.h }
func (c *BaseCell) SetWidth(w int)  { c.w = w }
func (c *BaseCell) SetHeight(h int) { c.h = h }
func (c *BaseCell) Add(in ...string) {
	for idx, val := range in {
		in[idx] = swellFont(val)
		if w := realLength(in[idx]); w > c.w {
			c.w = w
		}
	}
	c.val = append(c.val, in...)
	c.h += len(in)
}
func (c *BaseCell) Line(idx int) string {
	if idx >= len(c.val) || idx < 0 {
		return c.align.Repeat("", c.w)
	}
	return c.align.Repeat(c.val[idx], c.w)
}
func (c *BaseCell) Lines() (out []string) {
	for i := 0; i < c.h; i++ {
		out = append(out, c.Line(i))
	}
	return
}

type EmptyCell struct{ BaseCell }

func NewEmptyCell(w, h int) *EmptyCell {
	return &EmptyCell{
		BaseCell{
			w: w,
			h: h,
		},
	}
}
func (c *EmptyCell) Add(in ...string) { c.h += len(in) }
func (c *EmptyCell) Line(int) string  { return strings.Repeat(" ", c.w) }
func (c *EmptyCell) Lines() (out []string) {
	for i := 0; i < c.h; i++ {
		out = append(out, c.Line(i))
	}
	return
}
