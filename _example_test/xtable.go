package main

import (
	"fmt"

	"github.com/guojia99/go-tables/_example_test/zen"
	"github.com/guojia99/go-tables/tables"
)

type tt struct {
	ss    string
	No    int `json:"no" table:"zen no"`
	En    string
	Zn    string
	Slide []int
}

func structTableExample() {
	var v []tt
	for idx, val := range zen.List {
		v = append(v, tt{
			ss:    val,
			No:    idx,
			En:    val,
			Zn:    zen.ListZn[idx],
			Slide: []int{idx, idx * 2, idx ^ 2},
		})
	}

	opt := tables_back.NewOption()
	opt.Align = tables_back.AlignCenter

	t, err := tables_back.NewTable(v, opt)
	if err != nil {
		fmt.Println("[error]", err)
		return
	}
	fmt.Println(t.String())
}

func slideTableExample() {
	opt := tables_back.NewOption()
	opt.Align = tables_back.AlignRight

	t, err := tables_back.NewTable(zen.ListZn, opt)
	if err != nil {
		fmt.Println("[error]", err)
		return
	}
	fmt.Println(t.String())
}

func main() {
	structTableExample()
	slideTableExample()
}
