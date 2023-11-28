package main

import (
	"github.com/guojia99/go-tables/table"
)

func main() {
	t := tables.NewTable()

	opt := tables.Option{
		//AutoWidth: true,
		//AutoHeight: true,
		OrgPoint: tables.Address{
			Row: 3,
			Col: 4,
		},
		EndPoint: tables.Address{},
	}

	for i := 0; i < 6; i++ {
		var cells []tables.Cell
		for j := 0; j < 5; j++ {
			cells = append(cells, tables.NewCell(i*j, "中文123abc中文"))
		}
		t.AddBody(cells...)
	}
	t.SetOption(opt)
	t.String()
}
