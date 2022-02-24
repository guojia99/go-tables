package tables

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

type align int

const (
	AlignNone align = iota
	AlignLeft
	AlignCenter
	AlignRight
)

type Contour struct {
	H  string // "─"  Horizontal
	V  string // "│"  Vertical
	VH string // "┼"  Vertical Horizontal
	HU string // "┴"  Horizontal up
	HD string // "┬"  Horizontal down
	VL string // "┤"  Vertical left
	VR string // "├"  Vertical right
	UR string // "┐"  Up right
	UL string // "┌"  Up left
	DR string // "┘"  Down right
	DL string // "└"  Down left
}

var DefaultContour = &Contour{
	H: "─", V: "│", VH: "┼", HU: "┴", HD: "┬", VL: "┤", VR: "├", UR: "┐", UL: "┌", DR: "┘", DL: "└",
}
var dc = DefaultContour

func (c Contour) Header(vLen int) string {
	return fmt.Sprintf("%s%s%s\n", c.UL, strings.Repeat(c.H, vLen), c.UR)
}

func (c Contour) Center(vLen int) string {
	return fmt.Sprintf("")
}

func (c Contour) Footer(vLen int) string {
	return fmt.Sprintf("%s%s%s\n", c.DL, strings.Repeat(c.H, vLen), c.DR)
}

func RealLength(in string) int {
	in = color.ClearCode(in)
	return StringLength([]rune(in))
}

type Chinese struct {
	from rune
	to   rune
}

var lt = []Chinese{{0x2E80, 0x9FD0}, {0xAC00, 0xD7A3}, {0xF900, 0xFACE}, {0xFE00, 0xFE6C}, {0xFF00, 0xFF60}, {0x20000, 0x2FA1D}}

func StringLength(r []rune) int {
	length := len(r)
lens:
	for _, v := range r {
		for _, c := range lt {
			if v >= c.from && v <= c.to {
				length++
				continue lens
			}
		}
	}
	return length
}
