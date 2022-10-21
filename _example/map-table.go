package main

import (
	"fmt"
	"github.com/guojia99/go-tables/table"
)

func mapTable1() {
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	data := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
	}
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func mapTable2() {
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	data := map[string]interface{}{
		"number":  1,
		"string":  "guojia",
		"float":   2.4,
		"slide":   []int{1, 2, 3, 4},
		"complex": complex(1, -1),
		"key6": struct {
			a string
		}{a: "123"},
	}
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	mapTable1()
	mapTable2()
}
