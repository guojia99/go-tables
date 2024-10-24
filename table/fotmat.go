/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"regexp"
	"runtime"
	"strings"
)

type FullWidth struct {
	From  rune
	To    rune
	Width int
}

var _boxFullWidthMap = make(map[rune]int, 4096)

func SetBoxFullWith(in ...FullWidth) {
	for _, box := range in {
		for i := box.From; i < box.To; i++ {
			_boxFullWidthMap[i] = box.Width
		}
	}
}

var noWinCodeExpr = regexp.MustCompile(`\033\[[\d;?]+m`)
var defaultBoxFullWidth = []FullWidth{
	// 中文 (Chinese)
	{0x2E80, 0x9FD0, 2}, {0xAC00, 0xD7A3, 2}, {0xF900, 0xFACE, 2},
	{0xFE00, 0xFE6C, 2}, {0xFF00, 0xFF60, 2}, {0x20000, 0x2FA1D, 2},
	{0x3006, 0x303F, 2},

	// 日文 (Japanese)
	{0x3040, 0x309F, 2}, {0x30A0, 0x30FF, 2}, {0x4E00, 0x9FFF, 2},
	{0x3400, 0x4DBF, 2}, {0x20000, 0x2A6DF, 2},

	// 韩文 (Korean)
	{0x1100, 0x11FF, 2}, {0x3130, 0x318F, 2}, {0xAC00, 0xD7A3, 2},
}

func init() {
	SetBoxFullWith(defaultBoxFullWidth...)
}

func CodeFullWidth(in rune) int {
	w, ok := _boxFullWidthMap[in]
	if ok {
		return w
	}
	return 1
}

func stringRealLength(r []rune) (length int) {
	length = len(r) // base length
	for _, val := range r {
		w, ok := _boxFullWidthMap[val]
		if !ok {
			continue
		}
		length += w - 1
	}

	return
}

// RealLength 控制台实际显示长度
func RealLength(in string) int {
	in = strings.Trim(in, "\000")
	in = strings.TrimSpace(in)
	switch runtime.GOOS {
	case `windows`:
		return stringRealLength([]rune(in))
	}
	return stringRealLength([]rune(noWinCodeExpr.ReplaceAllString(in, "")))
}

func SplitWithRealLength(in string, maxLength int) []string {
	var (
		out       []string
		runeCache = []rune(in)
	)

	if maxLength <= 0 {
		maxLength = 32
	}

	var lastLen = maxLength
	var lastStr = ""
	for i := 0; i < len(runeCache); i++ {
		c := runeCache[i]
		l := CodeFullWidth(c)

		if lastLen-l >= 0 {
			lastLen -= l
			lastStr += string(c)
			continue
		}

		out = append(out, lastStr)
		lastStr = ""
		lastLen = maxLength
		i-- // 这个循环并没有处理到， 所以需要返回去处理
	}

	if lastStr != "" {
		out = append(out, lastStr)
	}
	return out
}
