package tables

import (
	"fmt"
	"strings"
)

type Cell struct {
	val    []string
	mw, mh int
	ag     align
}

func NewCell(data []string, ag align) *Cell {
	mw := 0
	for _, val := range data {
		if l := RealLength(val); l > mw {
			mw = l
		}
	}
	return &Cell{
		val: data,
		mw:  mw,
		mh:  len(data),
		ag:  ag,
	}
}

func (x *Cell) Width() int       { return x.mw }
func (x *Cell) Height() int      { return x.mh }
func (x *Cell) SetWidth(in int)  { x.mw = in }
func (x *Cell) SetHeight(in int) { x.mh = in }

func (x *Cell) Lines() []string {
	var out []string
	for idx := range x.val {
		out = append(out, x.Line(idx))
	}
	return out
}

func (x *Cell) Line(idx int) string {
	if idx >= len(x.val) {
		return x.serializer("")
	}
	return x.serializer(x.val[idx])
}

func (x *Cell) serializer(in string) string {
	if in == "" {
		return strings.Repeat(" ", x.mw)
	}
	l := x.mw - RealLength(in)
	switch x.ag {
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

func (x *Cell) String() string {
	out := dc.Header(x.mw)
	for _, val := range x.Lines() {
		out += fmt.Sprintf("%s%s%s\n", dc.V, val, dc.V)
	}
	out += dc.Footer(x.mw)
	return out
}
