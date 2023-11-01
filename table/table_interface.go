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
		Color() color.Style
		SetColor(color.Style) Cell
		SetWordWrap(b bool) Cell
		SetColWidth(w int)
		SetRowHeight(h int)
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
		SetColWidth(width int, cols ...int) error   // 设置某些列的宽
		SetRowHeight(height int, rows ...int) error // 设置某些行的高
	}

	TableUpdater interface {
		AddRows(Type RowType, cells ...Cell) (tb Table)                  // AddRows 添加一行数据
		SetRows(Type RowType, row int, cells ...Cell) (tb Table)         // SetRows 替换某一行数据
		InsertRows(Type RowType, startRow int, cells ...Cell) (tb Table) // InsertRows 插入一行数据
		SetCell(Type RowType, address Address, cell Cell) (tb Table)     // SetCell 替换某个单元格的数据
		DeleteRows(Type RowType, row int) (deleteCell Cells, ok bool)    // DeleteRows 删除某行数据
		AddBody(cells ...Cell) (tb Table)                                // AddBody 给主体添加一行
		AddHeader(cell ...Cell) (tb Table)                               // AddHeader 给头加一行
		AddFoots(cell ...Cell) (tb Table)                                // AddFoots 给脚注加一行
	}

	TableReader interface {
		AtRow(row int) (Cells, error)                                                 // AtRow 读取某一行
		AtCell(address Address) (cell Cell, ok bool)                                  // AtCell 读取某一格
		SortByCol(col int, less func(i, j interface{}) bool) (newTable Table)         // SortByCol 对某一列进行排序
		FilterByCol(col int, less func(interface{}) bool) (newTable Table)            // FilterByCol 对某一列进行过滤
		SearchCell(eq func(interface{}) bool) (cell Cell, address Address, err error) // SearchCell 搜索首个所需的单元格
	}

	TableColor interface {
		SetCellColor(address Address, color color.Color) error // SetCellColor 给某一单元格设置颜色
		SetCellColorByRow(row int, color color.Color) error    // SetCellColorByRow 给某一行设置颜色
		SetCellColorByCol(col int, color color.Color) error    // SetCellColorByCol 给某一列设置颜色
	}
)

type Table interface {
	Render(io.Writer) error // 渲染
	fmt.Stringer            // 直接把渲染结果输出

	TableArea
	TableColor
	TableReader
	TableUpdater
}

var _ Table = &UnimplementedTable{}

type UnimplementedTable struct{}

func (u UnimplementedTable) Render(writer io.Writer) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) String() string {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) OutputRect() (rect Address) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetOutputRect(rect Address) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetColWidth(width int, cols ...int) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetRowHeight(height int, rows ...int) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetCellColor(address Address, color color.Color) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetCellColorByRow(row int, color color.Color) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetCellColorByCol(col int, color color.Color) error {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AtRow(row int) (Cells, error) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AtCell(address Address) (cell Cell, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SortByCol(col int, less func(i interface{}, j interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) FilterByCol(col int, less func(interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SearchCell(eq func(interface{}) bool) (cell Cell, address Address, err error) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AddRows(Type RowType, cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetRows(Type RowType, row int, cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) InsertRows(Type RowType, startRow int, cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) SetCell(Type RowType, address Address, cell Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) DeleteRows(Type RowType, row int) (deleteCell Cells, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AddBody(cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AddHeader(cell ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (u UnimplementedTable) AddFoots(cell ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}
