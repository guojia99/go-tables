package main

import (
	"fmt"
	"github.com/guojia99/go-tables/table"
)

func slice2DTable1() {
	var opt = &table.Option{
		Contour: table.DefaultContour,
		Align:   table.AlignCenter,
	}
	data := [][]string{
		{"DATA1", "DATA2", "DATA3"},
		{"DATA4", "DATA5", "DATA6"},
		{"DATA7", "DATA8", "DATA9"},
		{"DATA10", "DATA11", "DATA12"},
	}
	tb, _ := table.SimpleTable(data, opt)
	fmt.Println(tb)
}

func main() {
	slice2DTable1()
}
