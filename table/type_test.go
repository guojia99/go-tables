package table

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
		want kind
	}{
		{
			name: "none",
			args: args{in: "none"},
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
			if got := parsingType(tt.args.in); got != tt.want {
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
			name: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHeadCapitalLetters(tt.args.in); got != tt.want {
				t.Errorf("isHeadCapitalLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
