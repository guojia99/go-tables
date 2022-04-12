package main

import (
	"fmt"
	"github.com/guojia99/go-tables/table"
)

func structSliceTable1() {
	data := []*struct {
		Str    string
		Val    string `table:"value"`
		Num    int    `json:"number"`
		NoUse  string `json:"-"`
		NoUse2 string `table:"-"`
	}{
		{"data1", "val1", 1, "no1", "no2"},
		{"data2", "val2", 2, "no2", "no3"},
		{"data3", "val3", 3, "no3", "no4"},
		{"data4", "val4", 4, "no4", "no5"},
		{"data5", "val5", 5, "no5", "no6"},
	}
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	tb, err := table.SimpleTable(data, opt)
	if err != nil {
		panic(err)
	}
	fmt.Println(tb)
}

func main() {
	structSliceTable1()
}
