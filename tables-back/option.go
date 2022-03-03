package tables_back

import "time"

type timeFormatEngine func(time.Time) string
type colorFormatEngine func(interface{}) string

type Option struct {
	Align       align
	Contour     Contour
	ColorEngine colorFormatEngine
	TimeEngine  timeFormatEngine
}

func (o *Option) NewEmptyTable() *Table {
	return emptyTable(o)
}

func NewOption() *Option {
	return &Option{
		Align:       AlignLeft,
		Contour:     DefaultContour,
		ColorEngine: DefaultColorStyles,
		TimeEngine:  DefaultSerializationTime,
	}
}
