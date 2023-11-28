/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"testing"
	"time"

	"github.com/gookit/color"

	color2 "github.com/guojia99/go-tables/table/color"
)

func TestNewCell(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
	}{
		{
			name: "one line",
			cell: NewCell("1111"),
		},
		{
			name: "one line word wrap",
			cell: NewCell("1111\n2222"),
		},
		{
			name: "fmt.Stringer",
			cell: NewCell(time.Now(), time.Now()),
		},
		{
			name: "many lines",
			cell: NewCell("1111", "2222", "3333"),
		},
		{
			name: "number",
			cell: NewCell(1, 2, 3.444),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fmt.Println(tt.cell)
			},
		)
	}
}

func TestBaseCell_UpdateColor(t *testing.T) {

	type args struct {
		style color.Style
	}
	tests := []struct {
		name string
		cell Cell
	}{
		{
			name: "red bg + font green",
			cell: NewCell("1234").SetColor(
				color.New(
					color.BgRed,
					color.Green,
				),
			),
		},
		{
			name: "hex red bg + hex green",
			cell: NewCell("5678").SetColor(
				color.New(
					color.Rgb(255, 0, 0, true).Color(),
					color.Rgb(0, 255, 0, false).Color(),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fmt.Println(tt.cell)
			},
		)
	}
}

func TestBaseCell_String(t *testing.T) {
	c := NewCell("1111")
	c.SetColor(color2.RedBgBlue)
	fmt.Println(c)
}

func TestBaseCell_Lines(t *testing.T) {
	c := NewCell("中文1234abc中文aaa", "第二行的数据, 不会被打扰")
	//c.SetColWidth(6)
	c.SetWordWrap(true)
	ls := c.Lines()
	for _, l := range ls {
		fmt.Printf("`%s`\n", l)
	}
}
