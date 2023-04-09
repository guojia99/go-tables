/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/4/3 下午7:09.
 *  * Author: guojia(https://github.com/guojia99)
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const format = `
/*
 * Copyright (c) 2023 gizwits.com All rights reserved.
 * Created: 2023/4/3 下午6:20.
 * Author: guojia(zjguo@gizwits.com)
 */

package tables

import "github.com/gookit/color"


// basic color
var (
	%s
)

// merge color
var (
	%s
)
`

	FgsKey := []string{
		"FgBlack", "FgRed", "FgGreen", "FgYellow", "FgBlue", "FgMagenta", "FgCyan", "FgWhite", "FgDefault", // basic
		"FgDarkGray", "FgLightRed", "FgLightGreen", "FgLightYellow", "FgLightBlue", "FgLightMagenta", "FgLightCyan",
		"FgLightWhite", "FgGray",
	}

	BgKey := []string{
		"BgBlack", "BgRed", "BgGreen", "BgYellow", "BgBlue", "BgMagenta", "BgCyan", "BgWhite", "BgDefault",
	}

	var (
		basicColorStr string
		mergeColorStr string
	)
	for _, key := range FgsKey {
		NewKey := strings.Replace(key, "Fg", "", -1)
		basicColorStr += fmt.Sprintf("\t%s = color.New(color.%s)\n", NewKey, key)

		for _, bg := range BgKey {
			mergeColorStr += fmt.Sprintf("\t%s = color.New(color.%s, color.%s)\n", NewKey+bg, key, bg)
		}
	}

	_ = os.WriteFile("table/color_utils_tools.go", []byte(fmt.Sprintf(format, basicColorStr, mergeColorStr)), 0644)

}
