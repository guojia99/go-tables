package main

import (
	"fmt"
	"github.com/guojia99/go-tables/tables"
)

type tt struct {
	aa string
	bb int
	cc string
}

func main() {
	v := []tt{
		{aa: "a1", bb: 1, cc: "c1"},
		{aa: "a2", bb: 2, cc: "c2"},
		{aa: "a3", bb: 3, cc: "c3"},
		{aa: "a4", bb: 4, cc: "c4"},
		{aa: "a5", bb: 5, cc: "c5"},
	}
	t, err := tables.NewTable(v, tables.NewTableOption())
	if err != nil {
		panic(err)
	}
	fmt.Println(t.String())
}
