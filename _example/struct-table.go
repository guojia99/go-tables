package main

import (
	"fmt"
	"github.com/guojia99/go-tables/table"
)

type structTable struct {
	Str    string
	Val    string `table:"value"`
	Num    int    `json:"number"`
	NoUse  string `json:"-"`
	NoUse2 string `table:"-"`
}

func structTable1() {
	data := structTable{
		Str:    "value",
		Val:    "val",
		Num:    111,
		NoUse:  "nouse",
		NoUse2: "nouse",
	}
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	structTable1()
}
