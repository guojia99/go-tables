package tables

import "C"
import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

type align int

const (
	AlignNone align = iota
	AlignLeft
	AlignCenter
	AlignRight
)

type Contour struct {
	H  string // '─'  Horizontal
	V  string // '│'  Vertical
	VH string // '┼'  Vertical Horizontal
	HU string // '┴'  Horizontal up
	HD string // '┬'  Horizontal down
	VL string // '┤'  Vertical left
	VR string // '├'  Vertical right
	UR string // '┐'  Up right
	UL string // '┌'  Up left
	DR string // '┘'  Down right
	DL string // '└'  Down left
}

var DefaultContour = &Contour{
	H: "─", V: "│", VH: "┼", HU: "┴", HD: "┬", VL: "┤", VR: "├", UR: "┐", UL: "┌", DR: "┘", DL: "└",
}
var dc = DefaultContour

func (c Contour) Header(vLen int) string {
	return fmt.Sprintf("%s%s%s\n", c.UL, strings.Repeat(c.H, vLen), c.UR)
}

func (c Contour) Center(vLen int) string {
	return fmt.Sprintf("%s%s%s", c.V, strings.Repeat(" ", vLen), c.V)
}

func (c Contour) Footer(vLen int) string {
	return fmt.Sprintf("%s%s%s\n", c.DL, strings.Repeat(c.H, vLen), c.DR)
}

func RealLength(in string) int {
	return stringLength([]rune(color.ClearCode(in)))
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

// You can add your full width interval
func AppTowBoxFonts(fws ...FullWidth) {
	fullWidth = append(fullWidth, fws...)
}
