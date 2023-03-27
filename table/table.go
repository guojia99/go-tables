/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package table

import (
	"fmt"
	"image"
)

type RowType int

const (
	Body RowType = iota
	Headers
	Foots
)

type Table interface {
	fmt.Stringer

	///*
	//	Rect and SetRect
	//	the rect is has cell numbers, for example:
	//		[] [] [] []
	//		[] [] [] []
	//	the rect is image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{4,2}}
	//	the rect is table output all lines
	//*/
	//Rect() (rect image.Rectangle)

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
	*/
	SetColWidth(width int, cols ...int)
	SetRowHeight(height int, rows ...int)

	// Page and Iterator
	Page(limit, offset int) (newTable Table)
	SetIterator(iterator Iterator)

	// Row CUDR
	SetHeader(cells ...interface{})
	SetFoots(cells ...interface{})
	AtCell(address image.Point) (cell Cell)
	SetCell(address image.Point, cell Cell)
	CreateRow(typ RowType, cells ...interface{}) (err error) // the interface it can also be a Cell
	UpdateRow(typ RowType, row int, cells ...interface{}) (err error)
	DeleteRow(typ RowType, row int) (err error)
	ReadRow(typ RowType, row int) (cells []Cell, err error)

	// Utils
	Sort(col int, less func(i, j int) bool) (newTable Table)
	Filter(col int, less func(interface{}) bool) (newTable Table)
	SearchCell(eq interface{}, str string) (cell Cell, address image.Point, err error)
	SearchRow(eq interface{}, str string) (cells []Cell, row int, err error)
}
