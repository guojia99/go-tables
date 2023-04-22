/*
 * Copyright (c) 2023 gizwits.com All rights reserved.
 * Created: 2023/4/14 下午5:46.
 * Author: guojia(zjguo@gizwits.com)
 */

package tables

import (
	`fmt`
	`reflect`
	`testing`
)

func _testNewCells(start, length int) []Cells {
	var out []Cells

	for i := start; i <= length; i++ {
		out = append(out, Cells{NewCell(i)})
	}
	return out
}

func TestBaseIterator(t *testing.T) {

	t.Run("jump page", func(t *testing.T) {
		iterator := NewBaseIterator(_testNewCells(1, 27))
		iterator.UpdateLimit(5)
		for i := 0; i < 7; i++ {
			got, _ := iterator.JumpPage(i)
			var want []Cells
			for j := i*5 + 1; j < (i+1)*5+1; j++ {
				if j <= 27 {
					want = append(want, Cells{NewCell(j)})
				}
			}
			if i == 6 {
				want = []Cells{{NewCell(26)}, {NewCell(27)}}
			}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("JumpPage() = %v, want %v", got, want)
			}
		}
	})

	t.Run("PrevLine", func(t *testing.T) {
		iterator := NewBaseIterator(_testNewCells(1, 6))
		iterator.NextLine()
		data, _ := iterator.PrevLine()
		fmt.Println(data)
	})
}
