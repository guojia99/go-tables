/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"io"
	"sync"

	"github.com/gookit/color"
)

var _ Table = &table{}

func NewTable() Table {
	return &table{}
}

type table struct {
	sync.Mutex
	opt  Option
	body Cells2D
}

func (t *table) Render(writer io.Writer) error {
	return nil
}

func (t *table) String() string {
	t.Lock()
	defer t.Unlock()

	t.autoScaling()

	if len(t.body) == 0 {
		return ""
	}

	var width, height = t.getWidthHeight()
	fmt.Println(width, height)
	// 获取每行的宽

	return ""
}

func (t *table) Clone() Table {
	t.Lock()
	defer t.Unlock()

	newTable := &table{
		body: make([]Cells, len(t.body)),
		opt:  t.opt,
	}
	copy(newTable.body, t.body)
	return newTable
}

func (t *table) SetContour(contour Contour) (tb Table) { t.opt.Contour = contour; return t }
func (t *table) SetOutputRect(start Address, end Address) Table {
	t.opt.OrgPoint = start
	t.opt.EndPoint = end
	return t
}
func (t *table) SetRowHeight(height int, rows ...int) error {
	return t.doRowWithFn(func(cell Cell) { cell.SetRowHeight(height) }, rows)
}
func (t *table) SetColWidth(width int, cols ...int) error {
	return t.doColWithFn(func(cell Cell) { cell.SetColWidth(width) }, cols)
}
func (t *table) SetCellColor(address Address, c color.Style) error {
	return t.doAddressWithFn(func(cell Cell) { cell.SetColor(c) }, address)
}
func (t *table) SetCellColorByRow(color color.Style, rows ...int) error {
	return t.doRowWithFn(func(cell Cell) { cell.SetColor(color) }, rows)
}
func (t *table) SetCellColorByCol(color color.Style, cols ...int) error {
	return t.doColWithFn(func(cell Cell) { cell.SetColor(color) }, cols)
}
func (t *table) SetCellWordWrap(address Address, wrap bool) error {
	return t.doAddressWithFn(func(cell Cell) { cell.SetWordWrap(wrap) }, address)
}
func (t *table) SetCellWordWrapByRow(wrap bool, rows ...int) error {
	return t.doRowWithFn(func(cell Cell) { cell.SetWordWrap(wrap) }, rows)
}
func (t *table) SetCellWordWrapByCol(wrap bool, cols ...int) error {
	return t.doColWithFn(func(cell Cell) { cell.SetWordWrap(wrap) }, cols)
}

func (t *table) UpdateOption(opts ...OptionFn) (tb Table) {
	for _, opt := range opts {
		opt(t)
	}
	return t
}

func (t *table) SetOption(opt Option) (tb Table) { t.opt = opt; return t }

func (t *table) AtRow(row int) (Cells, error) {
	var out Cells
	err := t.doRowWithFn(func(cell Cell) { out = append(out, cell) }, []int{row})
	return out, err
}

func (t *table) AtCell(address Address) (cell Cell, ok bool) {
	var out Cell
	ok = t.doAddressWithFn(func(cell Cell) { out = cell }, address) == nil
	return out, ok
}

func (t *table) SortByCol(col int, less func(i interface{}, j interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) FilterByCol(col int, less func(interface{}) bool) (newTable Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SearchCell(eq func(interface{}) bool) (cell Cell, address Address, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetRows(row int, cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) InsertRows(startRow int, cells ...Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) SetCell(address Address, cell Cell) (tb Table) {
	//TODO implement me
	panic("implement me")
}

func (t *table) InsertCols(startCol int, cells ...Cell) (tb Table) {
	t.Lock()
	defer t.Unlock()

	t.autoScaling()
	if startCol < 0 {
		startCol = 0
	}

	// 空的表格
	if len(t.body) == 0 {
		t.body = make(Cells2D, len(cells))
		for row, cell := range cells {
			t.body[row] = append(t.body[row], cell)
		}
		return tb
	}

	// 在前面的
	if startCol == 0 {

	}

	// 在中间的
	// 在末尾的
	// 超出范围的表格

	return t
}

func (t *table) DeleteRow(row int) (deleteCell Cells, err error) {
	t.Lock()
	defer t.Unlock()

	if row > len(t.body) || row < 0 {
		return nil, ErrRange
	}

	deleteCell = t.body[row]
	t.body = append(t.body[:row], t.body[row+1:]...)
	t.autoScaling()
	return
}

func (t *table) DeleteCol(col int) (deleteCell Cells, err error) {
	t.Lock()
	defer t.Unlock()

	t.autoScaling()
	if len(t.body) == 0 || len(t.body[0]) == 0 {
		return nil, nil
	}

	if col > len(t.body[0]) {
		return nil, ErrRange
	}

	for row := 0; row < len(t.body); row++ {
		deleteCell = append(deleteCell, t.body[row][col])
		t.body[row] = append(t.body[row][:col], t.body[row][col+1:]...) // remove
	}
	return
}

func (t *table) AddBody(cells ...Cell) (tb Table) {
	t.Lock()
	defer t.Unlock()

	t.body = append(t.body, cells)
	t.autoScaling()
	return t
}
