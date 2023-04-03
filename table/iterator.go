/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
 */

package table

type Iterator interface {
	// Size the Iterator max length
	Size() int
	// First the idx set to 0
	First()
	// Last the idx set to size
	Last()
	// Next in the next row of the table.
	Next() Cells
	// Prev from the previous row of the table.
	Prev() Cells
}

func NewSimpleIterator(data []Cells) *SimpleIterator {
	return &SimpleIterator{data: data, size: len(data)}
}

type SimpleIterator struct {
	// the size is max len
	size int
	idx  int
	data []Cells
}

func (s *SimpleIterator) Size() int { return s.size }
func (s *SimpleIterator) First()    { s.idx = 0 }
func (s *SimpleIterator) Last()     { s.idx = s.size }
func (s *SimpleIterator) Next() Cells {
	if s.idx < s.size-1 {
		s.idx += 1
		return s.data[s.idx]
	}
	return nil
}
func (s *SimpleIterator) Prev() Cells {
	if s.idx > 0 {
		s.idx -= 1
		return s.data[s.idx]
	}
	return nil
}
