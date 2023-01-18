package table

import (
	"fmt"
	"image"
	"io"

	"github.com/gookit/color"
)

type RowType int

const (
	Headers RowType = iota
	Body
	Foots
)

type Cell interface {
	fmt.Stringer
	Set(interface{})
	UpdateSize(w, h int)
	UpdateColor(font, bg color.Color256)
	SetWordWrap(b bool)
}

type Table interface {
	fmt.Stringer
	io.Writer
	image.Image

	// Base Configs
	Copy() (newTable Table)
	Rect() (rect image.Rectangle)
	SetRect(rect image.Rectangle) Table
	AtCell(address image.Point) (cell Cell)
	SetCell(address image.Point, cell Cell) Table

	// CUDR
	AddRow(typ RowType, cells ...interface{}) Table // the interface it can also be a Cell
	DeleteRow(typ RowType, row int) Table
	UpdateRow(typ RowType, row int, cells ...interface{}) Table
	ReadRow(typ RowType, row int) (cells []Cell)

	// Utils
	Sort(column int, less func(i, j int) bool) (newTable Table)
	SearchCell(eq interface{}, str string) (cell Cell, address image.Point, err error)
	SearchRow(eq interface{}, str string) (cells []Cell, row int, err error)
	Page(limit, offset int) (newTable Table)
}
