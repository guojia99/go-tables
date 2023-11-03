/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type RowType int

const (
	Body RowType = iota
	Headers
	Foots
)

type (
	Cell interface {
		fmt.Stringer
		Add(...string)
		Lines() []string
		IsEmpty() bool
		Color() color.Style
		SetColor(color.Style) Cell
		SetWordWrap(b bool) Cell
		SetColWidth(w int) Cell
		SetRowHeight(h int) Cell
		ColWidth() int
		RowHeight() int
	}
	Cells   []Cell
	Cells2D []Cells
)

type Address struct {
	Row int // 行
	Col int // 列
}

type (
	TableArea interface {
		OutputRect() (rect Address)                 // 获取输出的范围
		SetOutputRect(rect Address)                 // 设置输出的范围
		SetRowHeight(height int, rows ...int) error // 设置某些行的高,row从0开始算, 如果rows为空,则所有高都设置
		SetColWidth(width int, cols ...int) error   // 设置某些列的宽,col从0开始算, 如果cols为空,则所有的宽都设置
	}

	TableUpdater interface {
		SetRows(row int, cells ...Cell) (tb Table)         // SetRows 替换某一行数据
		InsertRows(startRow int, cells ...Cell) (tb Table) // InsertRows 插入一行数据
		InsertCols(startCol int, cells ...Cell) (tb Table) // InsertCols 插入一列数据
		SetCell(address Address, cell Cell) (tb Table)     // SetCell 替换某个单元格的数据
		DeleteRow(row int) (deleteCell Cells, err error)   // DeleteRows 删除某行数据
		DeleteCol(col int) (deleteCell Cells, err error)   // DeleteCols 删除某列
		AddBody(cells ...Cell) (tb Table)                  // AddBody 给主体添加一行
	}

	TableReader interface {
		AtRow(row int) (Cells, error)                                                 // AtRow 读取某一行
		AtCell(address Address) (cell Cell, ok bool)                                  // AtCell 读取某一格
		SortByCol(col int, less func(i, j interface{}) bool) (newTable Table)         // SortByCol 对某一列进行排序
		FilterByCol(col int, less func(interface{}) bool) (newTable Table)            // FilterByCol 对某一列进行过滤
		SearchCell(eq func(interface{}) bool) (cell Cell, address Address, err error) // SearchCell 搜索首个所需的单元格
	}

	TableCellUpdater interface {
		SetCellColor(address Address, color color.Style) error  // SetCellColor 给某一单元格设置颜色
		SetCellColorByRow(color color.Style, rows ...int) error // SetCellColorByRow 给某一行设置颜色
		SetCellColorByCol(color color.Style, cols ...int) error // SetCellColorByCol 给某一列设置颜色

		SetCellWordWrap(address Address, wrap bool) error  // SetCellWordWrap 给某一单元格设置换行
		SetCellWordWrapByRow(wrap bool, rows ...int) error // SetCellWordWrapByRow 给某一行设置换行
		SetCellWordWrapByCol(wrap bool, cols ...int) error // SetCellWordWrapByCol 给某一列设置换行
	}
)

type Table interface {
	Render(io.Writer) error // 渲染
	fmt.Stringer            // 直接把渲染结果输出
	Clone() Table

	TableArea
	TableCellUpdater
	TableReader
	TableUpdater
}
