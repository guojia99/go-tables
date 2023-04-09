/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import "image"

var _ Table = &table{}

type table struct {
	// outArea is output the table message.
	// if your outArea is [0, 0] - [3, 3], but the table inArea is [0, 0] - [4, 4], the output table is 3x3 not 4x4
	outArea image.Rectangle
	// inArea is input the table message data, is origin result.
	inArea image.Rectangle

	headers []Cells
	footers []Cells

	// datas
	iteratorIdx         int
	iterator            Iterator
	page, limit, offset int
	body                []Cells
}

func (t table) String() string {
	//TODO implement me
	panic("implement me")
}

func (t table) SetColWidth(width int, cols ...int) {
	//TODO implement me
	panic("implement me")
}

func (t table) SetRowHeight(height int, rows ...int) {
	//TODO implement me
	panic("implement me")
}

func (t table) Page(limit, offset int) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t table) SetIterator(iterator Iterator) {
	//TODO implement me
	panic("implement me")
}

func (t table) SetHeader(cells ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (t table) SetFoots(cells ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (t table) AtCell(address image.Point) (cell Cell) {
	//TODO implement me
	panic("implement me")
}

func (t table) SetCell(address image.Point, cell Cell) {
	//TODO implement me
	panic("implement me")
}

func (t table) CreateRow(typ RowType, cells ...interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t table) UpdateRow(typ RowType, row int, cells ...interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t table) DeleteRow(typ RowType, row int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t table) ReadRow(typ RowType, row int) (cells []Cell, err error) {
	//TODO implement me
	panic("implement me")
}

func (t table) Sort(col int, less func(i int, j int) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t table) Filter(col int, less func(interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t table) SearchCell(eq interface{}, str string) (cell Cell, address image.Point, err error) {
	//TODO implement me
	panic("implement me")
}

func (t table) SearchRow(eq interface{}, str string) (cells []Cell, row int, err error) {
	//TODO implement me
	panic("implement me")
}
