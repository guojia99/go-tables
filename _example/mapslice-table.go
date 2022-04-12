package main

import (
	"fmt"

	"github.com/guojia99/go-tables/table"
)

func mapSliceTable1() {
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	data := map[string][]string{
		"key1": {"value1", "value11", "value12", "value13"},
		"key2": {"value2", "value2", "value2", "value2"},
		"key3": {"value3", "value3", "value3", "value3"},
		"key4": {"value4", "value4", "value4", "value4"},
		"key5": {"value5", "value4", "value4", "value4"},
		"key6": {"value6", "value4", "value4", "value4"},
	}
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	mapSliceTable1()
}
