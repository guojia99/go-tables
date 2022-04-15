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
		"key1": {"key1-v", "key1-v", "key1-v", "key1-v", "key1-v", "key1-v"},
		"key2": {"key2-v2", "key2-v2", "key2-v2", "key2-v2"},
		"key3": {"key3-v3", "key3-v3", "key3-v3", "key3-v3"},
		"key4": {"key4-v4", "key4-v4", "key4-v4", "key4-v4"},
		"key5": {"key5-v5", "key5-v4", "key5-v4", "key5-v4"},
		"key6": {"key6-v6", "key6-v4", "key6-v4", "key6-v4"},
		"key7": {"key7-v6", "key7-v4", "key7-v4"},
	}
	fmt.Println()
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	mapSliceTable1()
}
