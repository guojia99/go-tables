package tables

import "time"

type timeEngineFormat func(time.Time) string
type colorStyle func(interface{}) string
type colorStyles []colorStyle

type Option struct {
	// engine list
	Align      align
	Contour    *Contour
	ColorStyle colorStyles
	TimeEngine timeEngineFormat
}

func NewOption() *Option {
	return &Option{
		Align:   AlignLeft,
		Contour: DefaultContour,
		ColorStyle: []colorStyle{
			DefaultColorStyles,
		},
		TimeEngine: DefaultSerializationTime,
	}
}
