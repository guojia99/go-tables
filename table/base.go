package table

//func NewTable() Table { return &table{} }

type table struct {
	page          int
	limit, offset int
	iterator      Iterator

	Headers [][]Cell
	Body    [][]Cell
	Footers [][]Cell
}
