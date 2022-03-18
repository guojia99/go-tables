package main

import (
	"fmt"
	"github.com/guojia99/go-tables/_example/zen"
	"github.com/guojia99/go-tables/table"
)

func sliceTable1() {
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	data := zen.ListZn
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)

	data = zen.List
	tb, _ = table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	sliceTable1()
}
