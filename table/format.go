package table

import (
	"strings"
)

var (
	// DefaultContour default
	DefaultContour = Contour{
		T: "─", R: "|",
		L: "|", D: "─",
		TL: "┌", TR: "┐",
		DL: "└", DR: "┘",
		TI: "┬", DI: "┴",
		LI: "├", RI: "┤",
		I: "┼", CH: "|", CV: "-",
	}

	// EmptyContour is empty
	EmptyContour = Contour{
		T: " ", R: " ",
		L: " ", D: " ",
		TL: " ", TR: " ",
		DL: " ", DR: " ",
		TI: " ", DI: " ",
		LI: " ", RI: " ",
		I: " ", CH: " ", CV: " ",
	}

	// MariaDBContour like to mysql
	MariaDBContour = Contour{
		T: "-", R: "|",
		L: "|", D: "-",
		TL: "+", TR: "+",
		DL: "+", DR: "+",
		TI: "+", DI: "+",
		LI: "+", RI: "+",
		I: "+", CH: "|", CV: "-",
	}
)

// Contour outlines are required when rendering the table
type Contour struct {
	// border like standard 8 azimuth
	//  ┌ ─ ┐       '┌' TL  '─' T  '┐' TR
	//  |   |   ->  '|' L          '|' R
	//  └ - ┘       '└  DL  '-' D  '┘' DR
	T, R, L, D     string
	TL, TR, DL, DR string
	// top | down intersection
	// for example, the two points highlighted above are the intersection points
	//  ┌ '┬' ─ ┐     '┬' TI -> top intersection
	//  |       | ->
	//  └ '┴' - ┘     '┴' DI -> down intersection
	TI, DI string
	// left | right intersection
	// for example, the two points highlighted above are the intersection points
	//   ┌ ─ ┐
	// 	'├'  |     '├' LI -> left intersection
	//   |  '┤'    '┤' RI -> right intersection
	//   └ - ┘
	LI, RI string
	// intersection location
	// ┌ ─ ┬ ─ ┬ - ┐
	// | - | - |'-'|  '-' CV -> center vertical
	// |  '|'  |   |  '|' CH -> center horizontal
	// ├ -'┼'- ┼ - ┤  '┼' I -> intersection
	// └ - ┴ - ┴ - ┘
	I, CH, CV string
}

func (c Contour) Handler(sw []int) string {
	out := c.TL
	for idx, s := range sw {
		out += strings.Repeat(c.T, s)
		if idx < len(sw)-1 {
			out += c.TI
		}
	}
	out += c.TR
	return out
}

func (c Contour) Footer(sw []int) string {
	out := c.DL
	for idx, s := range sw {
		out += strings.Repeat(c.D, s)
		if idx < len(sw)-1 {
			out += c.DI
		}
	}
	out += c.DR
	return out
}
