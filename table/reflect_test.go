/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"reflect"
	"testing"
)

func TestParsingTypeTBKind(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want TBKind
	}{
		{
			name: "struct",
			args: args{
				in: struct {
					name string
				}{name: "111"},
			},
			want: Struct,
		},
		{
			name: "struct slice",
			args: args{
				in: []struct {
					name string
				}{
					{name: "111"}, {name: "222"},
				},
			},
			want: StructSlice,
		},
		{
			name: "map",
			args: args{
				in: map[string]string{
					"aa": "aa",
				},
			},
			want: Map,
		},
		{
			name: "map slice",
			args: args{
				in: map[string][]int{
					"aa": {11223123, 312321, 32131},
				},
			},
			want: MapSlice,
		},
		{
			name: "array",
			args: args{
				in: [3]int{},
			},
			want: Slice,
		},
		{
			name: "slide",
			args: args{
				in: []int{1, 2, 3},
			},
			want: Slice,
		},
		{
			name: "slice 2d",
			args: args{
				in: [3][]int{},
			},
			want: Slice2D,
		},
		{
			name: "cell slice",
			args: args{
				in: []Cell{
					NewCell("1111"),
				},
			},
			want: CellSlice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsingTypeTBKind(tt.args.in); got != tt.want {
				t.Errorf("ParsingTypeTBKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseString(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "strings",
			args: args{
				in: "string",
			},
			wantErr: false,
		},
		{
			name: "fmt.Stringer",
			args: args{
				in: NewCell("111"),
			},
			wantErr: false,
		},
		{
			name: "not string",
			args: args{
				in: 1999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parseString(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_parseStruct(t *testing.T) {

	type testStruct struct {
		T1 string `json:"t1"`
		T2 int    `json:"t2"`
		T3 string `json:"-"`
	}
	tests := []struct {
		name       string
		in         interface{}
		wantHeader []Cell
		wantRow    []Cell
		wantErr    bool
	}{
		{
			name: "ok struct",
			in: testStruct{
				T1: "test1",
				T2: 1,
				T3: "test2",
			},
			wantHeader: []Cell{NewCell("t1"), NewCell("t2")},
			wantRow:    []Cell{NewCell("test1"), NewCell(1)},
			wantErr:    false,
		},
		{
			name:       "nil struct",
			in:         nil,
			wantHeader: nil,
			wantRow:    nil,
			wantErr:    true,
		},
		{
			name:       "error string input",
			in:         "string",
			wantHeader: nil,
			wantRow:    nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotRow, err := parseStruct(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("parseStruct() gotHeader = %+v, want %+v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotRow, tt.wantRow) {
				t.Errorf("parseStruct() gotRow = %+v, want %+v", gotRow, tt.wantRow)
			}
		})
	}
}

func Test_parseMapSlice(t *testing.T) {
	tests := []struct {
		name       string
		in         interface{}
		wantHeader Cells
		wantBody   Cells2D
		wantErr    bool
	}{
		{
			name: "map slide",
			in: map[string][]string{
				"test1": {"test1-value1", "test1-value2", "test1-value3"},
				"test3": {"test3-value1", "test3-value2", "test3-value3"},
				"test2": {"test2-value1"},
			},
			wantHeader: Cells{
				NewCell("test1"), NewCell("test2"), NewCell("test3"),
			},
			wantBody: Cells2D{
				{NewCell("test1-value1"), NewCell("test2-value1"), NewCell("test3-value1")},
				{NewCell("test1-value2"), NewEmptyCell(), NewCell("test3-value2")},
				{NewCell("test1-value3"), NewEmptyCell(), NewCell("test3-value3")},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotBody, err := parseMapSlice(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMapSlice() error = %+v, wantErr %+v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("parseMapSlice() gotHeader = %+v, want %+v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("parseMapSlice() gotBody = %+v, want %+v", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_parseStructSlice(t *testing.T) {
	type testStruct struct {
		T1 string `json:"t1"`
		T2 int    `json:"t2"`
		T3 string `json:"-"`
	}

	tests := []struct {
		name       string
		in         interface{}
		wantHeader Cells
		wantBody   Cells2D
		wantErr    bool
	}{
		{
			name: "struct slide",
			in: []testStruct{
				{T1: "1", T2: 2, T3: "3"},
				{T1: "4", T2: 5, T3: "6"},
				{T1: "", T2: 7, T3: "8"},
			},
			wantHeader: Cells{NewCell("t1"), NewCell("t2")},
			wantBody: Cells2D{
				{NewCell("1"), NewCell("2")},
				{NewCell("4"), NewCell("5")},
				{NewCell(""), NewCell("7")},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotBody, err := parseStructSlice(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStructSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("parseStructSlice() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("parseStructSlice() gotBody = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_parseSlice(t *testing.T) {

	tests := []struct {
		name     string
		in       interface{}
		wantBody Cells
		wantErr  bool
	}{
		{
			name: "slide",
			in:   []int{1, 2, 3},
			wantBody: Cells{
				NewCell("1"), NewCell("2"), NewCell("3"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, err := parseSlice(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("parseSlice() gotBody = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_parseSlice2D(t *testing.T) {
	tests := []struct {
		name     string
		in       interface{}
		wantBody Cells2D
		wantErr  bool
	}{
		{
			name: "slide 2d",
			in: [][]int{
				{1, 2, 3},
				{3, 2, 1},
			},
			wantBody: Cells2D{
				{NewCell("1"), NewCell("2"), NewCell("3")},
				{NewCell("3"), NewCell("2"), NewCell("1")},
			},
			wantErr: false,
		},
		{
			name: "array 2d",
			in: [3][4]int{
				{1, 2, 3, 4},
				{4, 3, 2, 1},
				{0, 1, 0, 2},
			},
			wantBody: Cells2D{
				{NewCell("1"), NewCell("2"), NewCell("3"), NewCell("4")},
				{NewCell("4"), NewCell("3"), NewCell("2"), NewCell("1")},
				{NewCell("0"), NewCell("1"), NewCell("0"), NewCell("2")},
			},
		},
		{
			name: "slide not neat",
			in: [][]int{
				{1, 2, 3},
				{1, 2, 3, 4},
				{1, 2},
				{1, 2, 3},
			},
			wantBody: Cells2D{
				{NewCell("1"), NewCell("2"), NewCell("3"), NewEmptyCell()},
				{NewCell("1"), NewCell("2"), NewCell("3"), NewCell("4")},
				{NewCell("1"), NewCell("2"), NewEmptyCell(), NewEmptyCell()},
				{NewCell("1"), NewCell("2"), NewCell("3"), NewEmptyCell()},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBody, err := parseSlice2D(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSlice2D() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("parseSlice2D() gotBody = %s, want %s", gotBody, tt.wantBody)
			}
		})
	}
}

func Test_parseMap(t *testing.T) {
	tests := []struct {
		name       string
		in         interface{}
		wantHeader Cells
		wantBody   Cells
		wantErr    bool
	}{
		{
			name: "map",
			in: map[int]int{
				1: 1, 2: 2, 3: 3,
			},
			wantHeader: Cells{NewCell("1"), NewCell("2"), NewCell("3")},
			wantBody:   Cells{NewCell("1"), NewCell("2"), NewCell("3")},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotBody, err := parseMap(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("parseMap() gotHeader = %s, want %s", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("parseMap() gotBody = %s, want %s", gotBody, tt.wantBody)
			}
		})
	}
}
