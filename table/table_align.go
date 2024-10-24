package tables

import "strings"

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
func (align Align) Repeat(in string, wantLen int) (out string) {
	if in == "" {
		return strings.Repeat(" ", wantLen)
	}
	realLen := RealLength(in)
	repeatLen := wantLen - realLen
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

func (align Align) Repeats(in []string, wantLen int) []string {
	var cache []string
	for _, val := range in {
		cache = append(cache, align.Repeat(val, wantLen))
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
