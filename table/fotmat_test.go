/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/24 下午10:01.
 * Author: guojia(zjguo@gizwits.com)
 */

package table

import (
	"fmt"
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
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := tt.align.Repeat(tt.args.in, uint(tt.args.wantLen)); gotOut != tt.wantOut {
				t.Errorf("Repeat() = `%v`, want `%v`", gotOut, tt.wantOut)
			}
		})
	}
}
