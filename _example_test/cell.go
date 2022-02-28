package main

import (
	"fmt"

	"github.com/guojia99/go-tables/_example_test/zen"
	"github.com/guojia99/go-tables/tables"
)

func main() {
	fmt.Println(tables.NewCell(tables.AlignLeft, zen.List...))

	var cells tables.Cells
	cells = append(cells, tables.NewCell(tables.AlignCenter, zen.ListZn...))
	cells = append(cells, tables.NewCell(tables.AlignCenter, zen.List...))
	fmt.Println(cells.Parse(tables.NewOption(), []int{1, 10}))
}
