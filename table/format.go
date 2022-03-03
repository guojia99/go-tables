package table

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

var (
	// DefaultContour
	DefaultContour = Contour{
		T: "─", R: "|",
		L: "|", D: "─",
		TL: "┌", TR: "┐",
		DL: "└", DR: "┘",
		TI: "┬", DI: "┴",
		LI: "├", RI: "┤",
		I: "┼", CH: "|", CV: "-",
	}

	// EmptyContour
	EmptyContour = Contour{
		T: "", R: "",
		L: "", D: "",
		TL: "", TR: "",
		DL: "", DR: "",
		TI: "", DI: "",
		LI: "", RI: "",
		I: "", CH: "", CV: "",
	}

	// MariaDBContour
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
