/*
 *  * Copyright (c) 2023 guojia99 All rights reserved.
 *  * Created: 2023/2/26 下午5:22.
 *  * Author: guojia(https://github.com/guojia99)
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
