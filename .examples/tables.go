package main

import "github.com/guojia99/go-tables/table"

func main() {
	t := tables.NewTable()

	opt := tables.Option{
		AutoWidth: true,
		//AutoHeight: true,
	}

	t.AddBody(
		tables.NewCell("1111"), tables.NewCell("aaaa"), tables.NewCell("xxxx"), tables.NewCell("1111"), tables.NewCell("aaaa"), tables.NewCell("xxxx"), tables.NewCell("1111"), tables.NewCell("aaaa"), tables.NewCell("xxxx"), tables.NewCell("1111"), tables.NewCell("aaaa"), tables.NewCell("xxxx"),
		tables.NewCell("1111"), tables.NewCell("aaaa"), tables.NewCell("xxxx"),
	)
	t.AddBody(tables.NewCell("2222"), tables.NewCell("bbbb"))
	t.AddBody(tables.NewCell("3333"), tables.NewCell("cccc"))
	t.AddBody(tables.NewCell("4444"), tables.NewCell("dddd"))
	t.SetOption(opt)
	t.String()
}
