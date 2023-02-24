/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/24 下午10:01.
 * Author: guojia(zjguo@gizwits.com)
 */

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
