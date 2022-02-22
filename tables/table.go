package tables

import "time"

type Contour struct {
	H  int32 // '─'  Horizontal
	V  int32 // '│'  Vertical
	VH int32 // '┼'  Vertical Horizontal
	HU int32 // '┴'  Horizontal up
	HD int32 // '┬'  Horizontal down
	VL int32 // '┤'  Vertical left
	VR int32 // '├'  Vertical right
	UR int32 // '┐'  Up right
	UL int32 // '┌'  Up left
	DR int32 // '┘'  Down right
	DL int32 // '└'  Down left
}

var DefaultContour = &Contour{'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'}

type Table struct {
	HeaderColors ColorStyles
	DataColors   ColorStyles
	// db is contour quit use
	db *Contour
	// The time of the output table is the serialized time
	useTimeEngine  bool
	timeEngineFunc func(time time.Time) string
}

func NewTable() *Table {
	return &Table{
		db:             DefaultContour,
		useTimeEngine:  true,
		timeEngineFunc: DefaultSerializationTime,
	}
}

// SetContour
func (t *Table) SetContour(in *Contour) *Table {
	t.db = in
	return t
}
