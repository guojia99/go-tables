/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"image"
	"sync"
)

var _ Table = &table{}

type table struct {
	lock sync.Mutex

	// outArea is output the table message.
	// if your outArea is [0, 0] - [3, 3], but the table inArea is [0, 0] - [4, 4], the output *table is 3x3 not 4x4
	outArea image.Rectangle
	// inArea is input the table message data, is origin result.
	inArea image.Rectangle

	headers []Cells
	footers []Cells

	// body fields
	iteratorIdx         int
	iterator            Iterator
	page, limit, offset int
	body                []Cells
}

func (t *table) String() string {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetColWidth(width int, cols ...int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if len(cols) == 0 {
		return
	}

	if t.iterator == nil {
		if cols[0] == -1 {
			for i := 0; i < len(t.body); i++ {
				for j := 0; j < len(t.body[i]); j++ {
					t.body[i][j].SetColWidth(width)
				}
			}
			return
		}

		for _, col := range cols {
			for i := 0; i < len(t.body); i++ {
				if len(t.body[i]) > col && col >= 0 {
					t.body[i][col].SetColWidth(width)
				}
			}
		}
		return
	}
	// 这里如果用迭代器怎么办？
}

func (t *table) SetRowHeight(height int, rows ...int) {
	t.lock.Lock()
	defer t.lock.Unlock()

	if len(rows) == 0 {
		return
	}

	if t.iterator == nil {
		if rows[0] == -1 {
			for i := 0; i < len(t.body); i++ {
				for j := 0; j < len(t.body[i]); j++ {
					t.body[i][j].SetRowHeight(height)
				}
			}
			return
		}

		for _, row := range rows {
			if len(t.body) > row && row >= 0 {
				for i := 0; i < len(t.body[row]); i++ {
					t.body[row][i].SetRowHeight(height)
				}
			}
		}
		return
	}
	// 这里如果用迭代器怎么办？
}

func (t *table) Page(limit, offset int) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetIterator(iterator Iterator) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetHeader(cells ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetFoots(cells ...interface{}) {
	//TODO implement me
	panic("implement me")
}

func (t *table) AtCell(address image.Point) (cell Cell) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetCell(address image.Point, cell Cell) {
	//TODO implement me
	panic("implement me")
}

func (t *table) CreateRow(typ RowType, cells ...interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) UpdateRow(typ RowType, row int, cells ...interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) DeleteRow(typ RowType, row int) (err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) ReadRow(typ RowType, row int) (cells []Cell, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) Sort(col int, less func(i int, j int) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) Filter(col int, less func(interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SearchCell(eq interface{}, str string) (cell Cell, address image.Point, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SearchRow(eq interface{}, str string) (cells []Cell, row int, err error) {
	//TODO implement me
	panic("implement me")
}
