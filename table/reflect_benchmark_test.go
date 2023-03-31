/*
 * Copyright (c) 2023 gizwits.com All rights reserved.
 * Created: 2023/3/31 下午4:24.
 * Author: guojia(zjguo@gizwits.com)
 */

package table

import (
	`image/color`
	`testing`
	`time`
)

func BenchmarkParsingTypeTBKind(b *testing.B) {
	b.Run(String.String(), func(b *testing.B) {
		ts := time.Now().String()
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(Struct.String(), func(b *testing.B) {
		ts := color.RGBA{R: 1, G: 2, B: 3}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(StructSlice.String(), func(b *testing.B) {
		ts := []color.RGBA{
			{R: 1, G: 2, B: 3},
			{R: 1, G: 2, B: 3},
			{R: 1, G: 2, B: 3},
		}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(Slice.String(), func(b *testing.B) {
		ts := []int{1, 2, 3, 4, 5, 6}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(Slice2D.String(), func(b *testing.B) {
		ts := [][]int{
			{1, 2, 3, 4, 5, 6},
			{1, 2, 3, 4, 5, 6},
			{1, 2, 3, 4, 5, 6},
		}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(Map.String(), func(b *testing.B) {
		ts := map[string]string{
			"1": "123", "2": "223", "3": "333",
		}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
	b.Run(MapSlice.String(), func(b *testing.B) {
		ts := map[string][]string{
			"1": {"1", "2", "3"},
			"2": {"1", "2", "3"},
			"3": {"1", "2", "3"},
		}
		for i := 0; i < b.N; i++ {
			parsingTypeTBKind(ts)
		}
	})
}
