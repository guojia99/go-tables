/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"image/color"
	"testing"
	"time"

	"github.com/guojia99/go-tables/table/mock"
)

func BenchmarkParsingTypeTBKind(b *testing.B) {
	b.Run(
		String.String(), func(b *testing.B) {
			ts := time.Now().String()
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		Struct.String(), func(b *testing.B) {
			ts := color.RGBA{R: 1, G: 2, B: 3}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		StructSlice.String(), func(b *testing.B) {
			ts := []color.RGBA{
				{R: 1, G: 2, B: 3},
				{R: 1, G: 2, B: 3},
				{R: 1, G: 2, B: 3},
			}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		Slice.String(), func(b *testing.B) {
			ts := []int{1, 2, 3, 4, 5, 6}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		Slice2D.String(), func(b *testing.B) {
			ts := [][]int{
				{1, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6},
				{1, 2, 3, 4, 5, 6},
			}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		Map.String(), func(b *testing.B) {
			ts := map[string]string{
				"1": "123", "2": "223", "3": "333",
				"4": "123", "5": "223", "6": "333",
			}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
	b.Run(
		MapSlice.String(), func(b *testing.B) {
			ts := map[string][]string{
				"1": {"1", "2", "3"},
				"2": {"1", "2", "3"},
				"3": {"1", "2", "3"},
			}
			for i := 0; i < b.N; i++ {
				parsingTypeTBKind(ts)
			}
		},
	)
}

func Benchmark_parseString(b *testing.B) {
	b.Run(
		"normal_string", func(b *testing.B) {
			data := "abcdefghijklmn"
			for i := 0; i < b.N; i++ {
				_, _ = parseString(data)
			}
		},
	)
	b.Run(
		"long_string", func(b *testing.B) {
			b.StopTimer()
			data := func() string {
				out := ""
				for i := 0; i < 10000; i++ {
					out += "abc"
				}
				return out
			}()
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				_, _ = parseString(data)
			}
		},
	)
}

func Benchmark_parseStruct(b *testing.B) {
	b.Run(
		"not_struct", func(b *testing.B) {
			data := "123"
			for i := 0; i < b.N; i++ {
				_, _, _ = parseStruct(data)
			}
		},
	)

	b.Run(
		"normal_struct", func(b *testing.B) {
			data := color.RGBA{
				R: 128,
				G: 128,
				B: 128,
				A: 128,
			}
			for i := 0; i < b.N; i++ {
				_, _, _ = parseStruct(data)
			}
		},
	)

	b.Run(
		"long_f_struct", func(b *testing.B) {
			b.StopTimer()
			data := new(mock.TestStruct1)
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = parseStruct(data)
			}
		},
	)
}

func Benchmark_parseStructSlice(b *testing.B) {
	b.Run(
		"normal_slice", func(b *testing.B) {
			b.StopTimer()
			var data []color.RGBA
			for i := 0; i < 10; i++ {
				data = append(data, color.RGBA{})
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				_, _, _ = parseStructSlice(data)
			}
		},
	)

	b.Run(
		"long_slice", func(b *testing.B) {
			b.StopTimer()
			var data []color.RGBA
			for i := 0; i < 10000; i++ {
				data = append(data, color.RGBA{})
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				_, _, _ = parseStructSlice(data)
			}
		},
	)

	b.Run(
		"long_struct", func(b *testing.B) {
			b.StopTimer()
			var data []mock.TestStruct1
			for i := 0; i < 100; i++ {
				data = append(data, mock.TestStruct1{})
			}
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				_, _, _ = parseStructSlice(data)
			}
		},
	)
}

func Benchmark_parseSlice(b *testing.B) {
	b.Run(
		"normal_slice", func(b *testing.B) {
			b.StopTimer()
			var data [100]int
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				_, _ = parseSlice(data)
			}
		},
	)

	b.Run(
		"long_slice", func(b *testing.B) {
			b.StopTimer()
			var data [10000]int
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				_, _ = parseSlice(data)
			}
		},
	)
}
