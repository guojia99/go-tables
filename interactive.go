/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"image"
	"io"
)

type RowType int

const (
	Body RowType = iota
	Headers
	Foots
)

type TableArea interface {
	/*
		Rect and SetRect
		the rect is has cell numbers, for example:
			A B C D
			E F G H
		the rect is image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{4,2}}, the rect is table output all lines:
			A B C D
			E F G H
		but your rect is image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{3,2}}
			A B C
			E F G

		如上所示, 提供表示输出的的表格范围, 如果你的表格很大, 那么只会按照你的限制格子数去做最终输出.
	*/
	Rect() (rect image.Rectangle)
	SetRect(rect image.Rectangle)

	/*
		SetColWidth and SetRowHeight
		----------
		|  |  |  | < row
		|  |  |  |
		----------
		 ^ col
		the SetColWidth can change the displayable width of all grids in this column,
		similarly, SetRowHeight can change all heights of the column.
		[!] but if you input the cols or rows has `-1` will change all columns or rows.

		设置table的宽高, 这里的设置是指整个table这一列或者这一行的数据是以多少行展示.
		一行则占用一个单位的字体数据， 数据默认为命令行的大小。
	*/
	SetColWidth(width int, cols ...int)
	SetRowHeight(height int, rows ...int)
}

type TableUpdater interface {
}

type Table interface {
	fmt.Stringer
	image.RGBA64Image
	io.Writer
	io.Reader

	TableArea

	AddRows(Type RowType, cells ...interface{}) (tb Table)
	SetRows(Type RowType, row int, cells interface{}) (tb Table)
	InsertRows(Type RowType, startRow int, cells ...interface{}) (tb Table)
	SetCell(Type RowType, address image.Point, cell interface{}) (tb Table)
	DeleteRows(Type RowType, row int) (deleteCell Cells, ok bool)
	AtCell(address image.Point) (cell Cell, ok bool)
}

type Table2 interface {
	fmt.Stringer

	// Page and Iterator
	Page(limit, offset int) (newTable Table)
	SetIterator(iterator Iterator)

	SetHeader(cells ...interface{})
	SetFoots(cells ...interface{})
	AtCell(address image.Point) (cell Cell)
	SetCell(address image.Point, cell Cell)
	CreateRow(typ RowType, cells ...interface{}) (err error) // the interface it can also be a Cell
	UpdateRow(typ RowType, row int, cells ...interface{}) (err error)
	DeleteRow(typ RowType, row int) (err error)
	ReadRow(typ RowType, row int) (cells []Cell, err error)

	Sort(col int, less func(i, j int) bool) (newTable Table)
	Filter(col int, less func(interface{}) bool) (newTable Table)
	SearchCell(eq interface{}, str string) (cell Cell, address image.Point, err error)
	SearchRow(eq interface{}, str string) (cells []Cell, row int, err error)
}
