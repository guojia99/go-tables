package utils

import (
	"fmt"
	"regexp"
	"runtime"
)

func SwellFont(in string) string { return fmt.Sprintf("%s%s%s", " ", in, " ") }

var noWinCodeExpr = regexp.MustCompile(`\033\[[\d;?]+m`)

func RealLength(in string) int {
	if runtime.GOOS != `windows` {
		return stringLength([]rune(noWinCodeExpr.ReplaceAllString(in, "")))
	}
	return stringLength([]rune(in))
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
