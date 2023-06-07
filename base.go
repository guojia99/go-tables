/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/4/9 下午10:27.
 * Author:  guojia(https://github.com/guojia99)
 */

package tables

import (
	"image"
	"sync"
)

var _ Table = &table{}

type table struct {
	lock sync.Mutex

	// outArea is output the table message.
	// if your outArea is [0, 0] - [3, 3], but the table inArea is [0, 0] - [4, 4], the output *table is 3x3 not 4x4
	outArea image.Rectangle
	// inArea is input the table message data, is origin result.
	inArea image.Rectangle

	page, limit, offset int
	body                []Cells

	headers []Cells
	footers []Cells

	iteratorIdx int
}
