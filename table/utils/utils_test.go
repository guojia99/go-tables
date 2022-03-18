package utils

import "testing"

func TestMax(t *testing.T) {
	type args struct {
		in []T
	}
	tests := []struct {
		name string
		args args
		want T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.in...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
