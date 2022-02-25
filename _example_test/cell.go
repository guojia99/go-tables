package main

import (
	"fmt"

	"github.com/guojia99/go-tables/tables"
	"github.com/guojia99/go-tables/zen"
)

func main() {
	fmt.Println(tables.NewCell(tables.AlignLeft, zen.List...))

	var cells tables.Cells
	cells = append(cells, tables.NewCell(tables.AlignCenter, zen.ListZn...))
	cells = append(cells, tables.NewCell(tables.AlignCenter, zen.List...))
	fmt.Println(cells.Parse(tables.NewTableOption(), []int{1, 10}))
}
