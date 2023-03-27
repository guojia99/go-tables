/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package table

import (
	`reflect`
	`testing`
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
					{name: "111"},
					{name: "222"},
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
	type args struct {
		in interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantHeader []Cell
		wantRow    []Cell
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHeader, gotRow, err := parseStruct(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotHeader, tt.wantHeader) {
				t.Errorf("parseStruct() gotHeader = %v, want %v", gotHeader, tt.wantHeader)
			}
			if !reflect.DeepEqual(gotRow, tt.wantRow) {
				t.Errorf("parseStruct() gotRow = %v, want %v", gotRow, tt.wantRow)
			}
		})
	}
}
