package table

import `fmt`

type Iterator interface {
	// Size the Iterator max length
	Size() int
	// First the idx set to 0
	First()
	// Last the idx set to size
	Last()
	// Next in the next row of the table.
	Next() []Cell
	// Prev from the previous row of the table.
	Prev() []Cell
}

type SimpleIterator struct {
	// the size is max len
	size int
	idx  int
	data [][]Cell
}

func (s *SimpleIterator) Size() int { return s.size }
func (s *SimpleIterator) First()    { s.idx = 0 }
func (s *SimpleIterator) Last()     { s.idx = s.size }
func (s *SimpleIterator) Next() []Cell {
	if s.idx < s.size-1 {
		s.idx += 1
		return s.data[s.idx]
	}
	return nil
}
func (s *SimpleIterator) Prev() []Cell {
	if s.idx > 0 {
		s.idx -= 1
		return s.data[s.idx]
	}
	return nil
}

func NewSimpleIterator(lists interface{}) (Iterator, error) {
	typ := parsingTypeTBKind(lists)
	switch typ {
	case IteratorSlice:
		return lists.(Iterator), nil
	case CellSlice:
		val := lists.([]Cell)
		return &SimpleIterator{size: len(val), idx: 0, data: [][]Cell{val}}, nil
	case String:
		// todo
	case Struct:

	case StructSlice:
	case Slice:
	case Slice2D:
	case Map:
	case MapSlice:
	default:
	}
	return nil, fmt.Errorf("the data type cannot be converted to an iterator")
}
