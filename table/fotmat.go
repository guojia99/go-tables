/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/24 下午10:01.
 * Author: guojia(zjguo@gizwits.com)
 */

package table

import (
	"regexp"
	"runtime"
	"strings"
)

var noWinCodeExpr = regexp.MustCompile(`\033\[[\d;?]+m`)

func RealLength(in string) int {
	in = strings.Trim(in, "\000")
	switch runtime.GOOS {
	case `windows`:
		return stringLength([]rune(in))
	}
	return stringLength([]rune(noWinCodeExpr.ReplaceAllString(in, "")))
}

type FullWidth struct {
	from rune
	to   rune
}

var fullWidth = []FullWidth{
	// Chinese
	{0x2E80, 0x9FD0}, {0xAC00, 0xD7A3}, {0xF900, 0xFACE},
	{0xFE00, 0xFE6C}, {0xFF00, 0xFF60}, {0x20000, 0x2FA1D},
	{12286, 12351},
}

func stringLength(r []rune) (length int) {
	length = len(r)
re:
	for _, val := range r {
		for _, twoBox := range fullWidth {
			if val >= twoBox.from && val <= twoBox.to {
				length++
				continue re
			}
		}
	}
	return
}

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight

	AlignTopLeft
	AlignTopCenter
	AlignTopRight

	AlignBottomLeft
	AlignBottomCenter
	AlignBottomRight
)

// Repeat the repeat strings transform by align, and other processing.
// in: base string
// wantLen: want string length
func (align Align) Repeat(in string, wantLen uint) (out string) {
	if in == "" {
		return strings.Repeat(" ", int(wantLen))
	}
	realLen := RealLength(in)
	repeatLen := int(wantLen) - realLen
	if repeatLen < 0 {
		return in[:wantLen]
	}

	switch align {
	case AlignLeft, AlignTopLeft, AlignBottomLeft:
		return in + strings.Repeat(" ", repeatLen)
	case AlignCenter, AlignTopCenter, AlignBottomCenter:
		leftL := repeatLen / 2
		rightL := repeatLen - leftL
		return strings.Repeat(" ", leftL) + in + strings.Repeat(" ", rightL)
	case AlignRight, AlignTopRight, AlignBottomRight:
		return strings.Repeat(" ", repeatLen) + in
	default:
		return strings.Repeat(" ", realLen)
	}
}

func (align Align) Repeats(in []string, count uint) []string {
	var cache []string
	for _, val := range in {
		if RealLength(val) == 0 {
			continue
		}
		cache = append(cache, align.Repeat(val, count))
	}
	repeatLen := len(in) - len(cache)
	switch align {
	case AlignTopRight, AlignTopLeft, AlignTopCenter:
		return append(cache, make([]string, repeatLen)...)
	case AlignBottomRight, AlignBottomLeft, AlignBottomCenter:
		return append(make([]string, repeatLen), cache...)
	case AlignRight, AlignLeft, AlignCenter:
		leftL := repeatLen / 2
		rightL := repeatLen - leftL
		return append(make([]string, leftL), append(cache, make([]string, rightL)...)...)
	}
	return in
}

func isHeadCapitalLetters(in string) bool {
	if len(in) != 0 || !('A' <= in[0] && in[0] <= 'Z') {
		return false
	}
	return true
}
