package tables

import (
	"fmt"
	"strings"
)

type Cell struct {
	val    []string
	w, h   int
	mw, mh int
	ag     align
}

func NewCell(ag align, data ...string) *Cell {
	mw := 0
	for _, val := range data {
		if l := RealLength(val); l >= mw {
			mw = l
		}
	}
	return &Cell{
		val: data,
		w:   mw,
		h:   len(data),
		mw:  mw,
		mh:  len(data),
		ag:  ag,
	}
}

func (c *Cell) String() string {
	out := dc.Header(c.mw)
	for _, val := range c.Lines() {
		out += fmt.Sprintf("%s%s%s\n", dc.V, val, dc.V)
	}
	out += dc.Footer(c.mw)
	return out
}

func (c *Cell) Width() int     { return c.w }
func (c *Cell) Height() int    { return c.h }
func (c *Cell) MaxWidth() int  { return c.mw }
func (c *Cell) MaxHeight() int { return c.mh }

func (c *Cell) Lines() []string {
	var out []string
	for idx := range c.val {
		out = append(out, c.Line(idx))
	}
	return out
}

func (c *Cell) Line(idx int) string {
	if idx >= len(c.val) {
		return c.serializer("")
	}
	return c.serializer(c.val[idx])
}

func (c *Cell) serializer(in string) string {
	if in == "" {
		return strings.Repeat(" ", c.mw)
	}
	l := c.mw - RealLength(in)
	switch c.ag {
	case AlignLeft, AlignNone:
		in = in + strings.Repeat(" ", l)
	case AlignRight:
		in = strings.Repeat(" ", l) + in
	case AlignCenter:
		leftL := l / 2
		rightL := l - leftL
		in = strings.Repeat(" ", leftL) + in + strings.Repeat(" ", rightL)
	}
	return in
}

type Cells []*Cell

// Max output -> max Widths and max Height
func (cs Cells) MaxWidths() []int {
	mws := make([]int, len(cs))
	for idx, val := range cs {
		if val.mw >= mws[idx] {
			mws[idx] = val.mw
		}
	}
	return mws
}

func (cs Cells) MaxH() int {
	mh := 0
	for _, val := range cs {
		if val.mh > mh {
			mh = val.mh
		}
	}
	return mh
}

func (cs Cells) Parse(opt *TableOption, mws []int) string {
	out := ""
	if mws == nil || len(mws) != len(cs) {
		mws = cs.MaxWidths()
	}
	for idx, val := range cs {
		if val.mw >= mws[idx] {
			mws[idx] = val.mw
		}
	}
	mh := cs.MaxH()
	for idx, val := range cs {
		val.mw = mws[idx]
		val.mh = mh
	}
	for h := 0; h < mh; h++ {
		out += opt.Contour.V
		for _, val := range cs {
			out += val.Line(h) + opt.Contour.V
		}
		out += "\n"
	}
	return out
}
