package tables

import "time"

type Table struct {
	HeaderColors ColorStyles
	DataColors   ColorStyles
	// db is contour quit use
	Contour *Contour
	// The time of the output table is the serialized time
	useTimeEngine  bool
	timeEngineFunc func(time time.Time) string
}

func NewTable() *Table {
	return &Table{
		HeaderColors:   NewDefaultColorStyles(),
		DataColors:     NewDefaultColorStyles(),
		Contour:        DefaultContour,
		useTimeEngine:  true,
		timeEngineFunc: DefaultSerializationTime,
	}
}

// SetContour
func (t *Table) SetContour(in *Contour) *Table {
	t.Contour = in
	return t
}
