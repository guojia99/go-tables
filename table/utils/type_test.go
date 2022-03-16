package utils

import (
	"testing"
)

func Test_parsingType(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name string
		args args
		want Kind
	}{
		{
			name: "none string",
			args: args{in: "none"},
			want: None,
		},
		{
			name: "none int",
			args: args{in: 1},
			want: None,
		},
		{
			name: "none float",
			args: args{in: 1.1},
			want: None,
		},
		{
			name: "none complex",
			args: args{in: complex(1, 11)},
			want: None,
		},
		{
			name: "none nil",
			args: args{in: nil},
			want: None,
		},
		{
			name: "none func",
			args: args{in: func() {}},
			want: None,
		},
		{
			name: "struct",
			args: args{in: struct{}{}},
			want: Struct,
		},
		{
			name: "map",
			args: args{in: map[string]string{}},
			want: Map,
		},
		{
			name: "map slice",
			args: args{in: map[string][]string{}},
			want: MapSlice,
		},
		{
			name: "struct slice",
			args: args{in: []struct{}{}},
			want: StructSlice,
		},
		{
			name: "slice",
			args: args{in: []string{}},
			want: Slice,
		},
		{
			name: "slice 2D",
			args: args{in: [][]string{}},
			want: Slice2D,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsingType(tt.args.in); got != tt.want {
				t.Errorf("parsingType() = %v, want %v", got, tt.want)
			}
		})
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
			name: "OK",
			args: args{in: "ADDDs"},
			want: true,
		},
		{
			name: "none",
			args: args{in: "sDD"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHeadCapitalLetters(tt.args.in); got != tt.want {
				t.Errorf("isHeadCapitalLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
