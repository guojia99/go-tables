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
func (c Contour) Handler(sw []uint) string {
	out := c.TL
	for idx, s := range sw {
		out += strings.Repeat(c.T, int(s))
		if idx < len(sw)-1 {
			out += c.TI
		}
	}
	out += c.TR
	return out + "\n"
}

func (c Contour) Intersection(sw []uint) string {
	out := c.LI
	for idx, s := range sw {
		out += strings.Repeat(c.CV, int(s))
		if idx < len(sw)-1 {
			out += c.I
		}
	}
	out += c.RI
	return out + "\n"
}

func (c Contour) Footer(sw []uint) string {
	out := c.DL
	for idx, s := range sw {
		out += strings.Repeat(c.D, int(s))
		if idx < len(sw)-1 {
			out += c.DI
		}
	}
	out += c.DR
	return out + "\n"
}
