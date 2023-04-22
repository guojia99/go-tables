/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

func NewBaseIterator(data []Cells) Iterator {
	limit := len(data)
	if len(data) > 20 {
		limit = 20
	}
	return &BaseIterator{data: data, info: IteratorInfo{Count: len(data), Page: 1, Limit: limit, Line: 0}}
}

type (
	IteratorInfo struct {
		Count int
		Page  int // start 0
		Limit int // start 1
		Line  int // start 0
	}

	Iterator interface {
		// GetInfo the current info data.
		GetInfo() IteratorInfo
		// GetContent the data from the current row to the beginning to one page.
		GetContent() []Cells
		// UpdateLimit the length of one page, the minimum is 1.
		UpdateLimit(limit int) bool
		// JumpPage jump to a page.
		JumpPage(page int) ([]Cells, bool)
		// PrevLine jump the current line to the previous line and return the new current line.
		PrevLine() (Cells, bool)
		// NextLine jump the current line to the next line and return the new current line.
		NextLine() (Cells, bool)
	}
)

type BaseIterator struct {
	// the size is max len
	info IteratorInfo
	data []Cells
}

func (s *BaseIterator) GetInfo() IteratorInfo { return s.info }

func (s *BaseIterator) GetContent() []Cells {
	// Locate to which page the current row is
	maxPage := s.info.Count / s.info.Limit
	if s.info.Page >= maxPage {
		return s.data[s.info.Line:]
	}
	return s.data[s.info.Line : s.info.Line+s.info.Limit]
}

func (s *BaseIterator) JumpPage(page int) ([]Cells, bool) {
	// if cur page > max page
	if page <= 0 || page > s.info.Count/s.info.Limit {
		return s.GetContent(), false
	}
	s.info.Page = page
	s.info.Line = page * s.info.Limit
	return s.GetContent(), true
}

func (s *BaseIterator) UpdateLimit(limit int) bool {
	if limit <= 0 || limit > s.info.Count {
		return false
	}
	s.info.Limit = limit
	s.info.Page = (s.info.Line / s.info.Limit) + 1 // ceil
	return true
}

func (s *BaseIterator) PrevLine() (Cells, bool) {
	if s.info.Line == 0 { // the line start 1
		return nil, false
	}
	s.info.Line -= 1
	s.info.Page = (s.info.Line / s.info.Limit) + 1 // ceil
	return s.data[s.info.Line], true
}

func (s *BaseIterator) NextLine() (Cells, bool) {
	if s.info.Line+1 > s.info.Count {
		return nil, false
	}
	s.info.Line += 1
	s.info.Page = (s.info.Line / s.info.Limit) + 1 // ceil
	return s.data[s.info.Line], true
}
