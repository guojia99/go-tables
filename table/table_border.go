package tables

import "strings"

var (
	// DefaultContour default
	/*
		┌──────────┬──────┐
		|    A     |  B   |
		├──────────┼──────┤
		|   112    |  2   |
		|    1     | 223  |
		|    4     |  5   |
		| 72312312 | 2328 |
		└──────────┴──────┘
	*/
	DefaultContour = Contour{
		T: "─", R: "|",
		L: "|", D: "─",
		TL: "┌", TR: "┐",
		DL: "└", DR: "┘",
		TI: "┬", DI: "┴",
		LI: "├", RI: "┤",
		I: "┼", CH: "|", CV: "─",
	}

	// EmptyContour is empty
	/*
	     A        B

	    112       2
	     1       223
	     4        5
	  72312312   2328
	*/
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
	/*
		+----------+------+
		|    A     |  B   |
		+----------+------+
		|   112    |  2   |
		|    1     | 223  |
		|    4     |  5   |
		| 72312312 | 2328 |
		+----------+------+
	*/
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

// Handler output : ┌──────────┬──────┐
func (c Contour) Handler(sw []int) string {
	out := c.TL
	for idx, s := range sw {
		out += strings.Repeat(c.T, s)
		if idx < len(sw)-1 {
			out += c.TI
		}
	}
	out += c.TR + "\n"
	return out
}

// Intersection	output :	├──────────┼──────┤
func (c Contour) Intersection(sw []int) string {
	out := c.LI
	for idx, s := range sw {
		out += strings.Repeat(c.CV, s)
		if idx < len(sw)-1 {
			out += c.I
		}
	}
	out += c.RI + "\n"
	return out
}

// Content	output : `|  data1  |  data2   |`
func (c Contour) Content(values []string) string {

	out := c.L
	for idx, val := range values {
		out += val
		if idx < len(values)-1 {
			out += c.CH
		}
	}
	out += c.R + "\n"
	return out
}

// Footer output :	└──────────┴──────┘
func (c Contour) Footer(sw []int) string {
	out := c.DL
	for idx, s := range sw {
		out += strings.Repeat(c.D, s)
		if idx < len(sw)-1 {
			out += c.DI
		}
	}
	out += c.DR + "\n"
	return out
}
