package main

import (
	"fmt"

	"github.com/guojia99/go-tables/_example_test/zen"
	"github.com/guojia99/go-tables/tables"
)

func main() {
	fmt.Println(tables_back.NewCell(tables_back.AlignLeft, zen.List...))

	var cells tables_back.Cells
	cells = append(cells, tables_back.NewCell(tables_back.AlignCenter, zen.ListZn...))
	cells = append(cells, tables_back.NewCell(tables_back.AlignCenter, zen.List...))
	fmt.Println(cells.Parse(tables_back.NewOption(), []int{1, 10}))
}
