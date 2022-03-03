package tables_back

import (
	"fmt"
	"strings"
)

func swellFont(in string) string {
	return fmt.Sprintf("%s%s%s", " ", in, " ")
}

type Cell struct {
	val    []string
	w, h   int
	mw, mh int
	ag     align
}

// NewCell create a cell from a list of strings
func NewCell(ag align, data ...string) *Cell {
	mw := 0
	for idx, val := range data {
		data[idx] = swellFont(val)
		if l := RealLength(data[idx]); l >= mw {
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

// String output a single cell
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

// Lines get every row of this cell
func (c *Cell) Lines() []string {
	var out []string
	for idx := range c.val {
		out = append(out, c.Line(idx))
	}
	return out
}

// Line get a row of the cell
func (c *Cell) Line(idx int) string {
	if idx >= len(c.val) {
		return c.serializer("")
	}
	return c.serializer(c.val[idx])
}

// serializer adjust internal size and font orientation
func (c *Cell) serializer(in string) (out string) {
	if in == "" {
		return strings.Repeat(" ", c.mw)
	}
	l := c.mw - RealLength(in)
	switch c.ag {
	case AlignLeft, AlignNone:
		return in + strings.Repeat(" ", l)
	case AlignRight:
		return strings.Repeat(" ", l) + in
	case AlignCenter:
		leftL := l / 2
		rightL := l - leftL
		return strings.Repeat(" ", leftL) + in + strings.Repeat(" ", rightL)
	}
	return
}

func MaxH(cells []*Cell) int {
	mh := 0
	for _, val := range cells {
		if val.MaxHeight() > mh {
			mh = val.MaxHeight()
		}
	}
	return mh
}

func MaxCellsWidths(cells []*Cell) []int {
	mws := make([]int, len(cells))
	for idx, val := range cells {
		if val.MaxWidth() >= mws[idx] {
			mws[idx] = val.MaxWidth()
		}
	}
	return mws
}

type Cells []*Cell

func (cs Cells) Parse(mws []int, opt *Option) string {
	out := ""
	if mws == nil || len(mws) != len(cs) {
		mws = MaxCellsWidths(cs)
	}
	for idx, val := range cs {
		if val.MaxWidth() >= mws[idx] {
			mws[idx] = val.MaxWidth()
		}
	}
	mh := MaxH(cs)
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
