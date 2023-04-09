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
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.cell)
		})
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
			cell: NewCell("1234").SetColor(color.New(
				color.BgRed,
				color.Green,
			)),
		},
		{
			name: "hex red bg + hex green",
			cell: NewCell("5678").SetColor(color.New(
				color.Rgb(255, 0, 0, true).Color(),
				color.Rgb(0, 255, 0, false).Color(),
			)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.cell)
		})
	}
}

func TestBaseCell_String(t *testing.T) {
	c := NewCell("1111")
	c.SetColor(RedBgBlue)
	fmt.Println(c)
}
