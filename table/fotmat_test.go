/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestRealLength(t *testing.T) {
	b := []byte{0, 66, 68, 70, 0, 0, 50, 52, 54, 52, 55, 49, 0, 0, 0, 71, 0, 0, 0, 2, 0, 0, 0, 1,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(RealLength(string(b)))

	in := string(b)
	fmt.Println(RealLength(in))
}

func TestAlign_Repeat(t *testing.T) {

	baseStr := "base tests"
	baseLen := len(baseStr)

	type args struct {
		in      string
		wantLen int
	}
	tests := []struct {
		name    string
		align   Align
		args    args
		wantOut string
	}{
		{
			name:  "base",
			align: AlignLeft,
			args: args{
				in:      baseStr,
				wantLen: baseLen,
			},
			wantOut: "base tests",
		},
		{
			name:  "less required length",
			align: AlignLeft,
			args: args{
				in:      baseStr,
				wantLen: 6,
			},
			wantOut: "base t",
		},
		{
			name:  "move then required length",
			align: AlignTopLeft,
			args: args{
				in:      baseStr,
				wantLen: 13,
			},
			wantOut: baseStr + strings.Repeat(" ", 3),
		},
		{
			name:    "empty",
			align:   0,
			args:    args{},
			wantOut: "",
		},
		{
			name:  "right",
			align: AlignRight,
			args: args{
				in:      baseStr,
				wantLen: 15,
			},
			wantOut: strings.Repeat(" ", 5) + baseStr,
		},
		{
			name:  "center 1",
			align: AlignCenter,
			args: args{
				in:      baseStr,
				wantLen: 15,
			},
			wantOut: strings.Repeat(" ", 2) + baseStr + strings.Repeat(" ", 3),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotOut := tt.align.Repeat(tt.args.in, int(tt.args.wantLen)); gotOut != tt.wantOut {
					t.Errorf("Repeat() = `%v`, want `%v`", gotOut, tt.wantOut)
				}
			},
		)
	}
}

func Test_isHeadCapitalLetters(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := isHeadCapitalLetters(tt.args.in); got != tt.want {
					t.Errorf("isHeadCapitalLetters() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_stringLength(t *testing.T) {
	type args struct {
		r []rune
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{
			name: "中文",
			args: args{
				r: []rune("你好世界"),
			},
			wantLength: 8,
		},
		{
			name: "日文",
			args: args{
				r: []rune("こんにちは 世界"),
			},
			wantLength: 15,
		},
		{
			name: "韩文",
			args: args{
				r: []rune("안녕하세요 세계"),
			},
			wantLength: 15,
		},
		{
			name: "俄文",
			args: args{
				r: []rune("Привет мир"),
			},
			wantLength: 10,
		},
		{
			name: "英文",
			args: args{
				r: []rune("hello world"),
			},
			wantLength: 11,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if gotLength := stringRealLength(tt.args.r); gotLength != tt.wantLength {
					t.Errorf("stringLength() = %v, want %v", gotLength, tt.wantLength)
				}
			},
		)
	}
}

func TestSplitWithRealLength(t *testing.T) {
	type args struct {
		in        string
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "0",
			args: args{
				in:        "123456789中",
				maxLength: 15,
			},
			want: []string{
				"123456789中",
			},
		},
		{
			name: "1",
			args: args{
				in:        "123456789中文123456789中文",
				maxLength: 12,
			},
			want: []string{
				"123456789中", "文123456789", "中文",
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := SplitWithRealLength(tt.args.in, tt.args.maxLength); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("SplitWithRealLength() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func BenchmarkSplitWithRealLength(b *testing.B) {
	b.Run(
		"long msg", func(b *testing.B) {
			b.StopTimer()
			var data string
			for i := 0; i < 10000; i++ {
				data += "1"
			}
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				SplitWithRealLength(data, 100)
			}
		},
	)

	b.Run(
		"shout msg", func(b *testing.B) {
			b.StopTimer()
			var data = "123456789中文"
			b.StartTimer()

			for i := 0; i < b.N; i++ {
				SplitWithRealLength(data, 100)
			}
		},
	)
}
