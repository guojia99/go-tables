package tables

import (
	tty "github.com/mattn/go-tty"
)

// autoScaling 计算每一行最后一个empty的位置, 多出来的删除掉, 少的需要扩充
func (t *table) autoScaling() {
	var maxCol = 0

	// 从后往前找非empty的位置, 记录所有行最后一个非empty的长度
	for i := 0; i < len(t.body); i++ {
		for j := len(t.body[i]) - 1; j >= 0; j-- {
			if !t.body[i][j].IsEmpty() {
				if j >= maxCol {
					maxCol = j
				}
				break
			}
		}
	}
	maxCol += 1

	// 遍历每行
	for i := 0; i < len(t.body); i++ {
		length := len(t.body[i])
		if length == maxCol {
			continue
		}
		if length > maxCol {
			t.body[i] = append(t.body[i][:maxCol], t.body[i][maxCol+1:]...)
			continue
		}
		t.body[i] = append(t.body[i], NewEmptyCells(maxCol-length)...)
	}
}

func (t *table) doColWithFn(fn func(cell Cell), cols []int) error {
	t.Lock()
	defer t.Unlock()
	t.autoScaling()
	if len(cols) != 0 {
		for _, col := range cols {
			if col < 0 {
				return ErrRange
			}
		}
		for _, col := range cols {
			for row := 0; row < len(t.body); row++ {
				fn(t.body[row][col])
			}
		}
		return nil
	}
	if len(t.body) >= 1 {
		for col := 0; col < len(t.body[0]); col++ {
			for row := 0; row < len(t.body); row++ {
				fn(t.body[row][col])
			}
		}
	}

	return nil
}

func (t *table) doRowWithFn(fn func(cell Cell), rows []int) error {
	t.Lock()
	defer t.Unlock()

	if len(rows) == 0 {
		for i := 0; i < len(t.body); i++ {
			rows = append(rows, i)
		}
	}
	for _, row := range rows {
		if row < 0 || row >= len(t.body) {
			return ErrRange
		}
		for col := 0; col < len(t.body[row]); col++ {
			fn(t.body[row][col])
		}
	}
	return nil
}

func (t *table) doAddressWithFn(fn func(cell Cell), address Address) error {
	t.Lock()
	defer t.Unlock()
	t.autoScaling()
	if len(t.body) == 0 || len(t.body[0]) == 0 {
		return nil
	}

	if address.Row < 0 || address.Col < 0 {
		return ErrAddress
	}

	if address.Row >= len(t.body) || address.Col >= len(t.body[0]) {
		return ErrAddress
	}

	fn(t.body[address.Row][address.Col])
	return nil
}

func (t *table) getWidthHeight() (int, int) {
	if len(t.body) == 0 {
		return 0, 0
	}

	// 重新计算长度
	org := t.opt.OrgPoint
	end := t.opt.EndPoint
	if end.Col == t.opt.OrgPoint.Col && end.Col == 0 {
		org.Col, end.Col = 0, len(t.body[0])
	}
	if end.Row == org.Row && end.Row == 0 {
		org.Row, end.Row = 0, len(t.body)
	}

	// 正常情况
	w := end.Col - org.Col
	h := end.Row - org.Row
	if !t.opt.AutoWidth && !t.opt.AutoHeight {
		return w, h
	}

	// 依据屏幕的情况
	ttyIn, err := tty.Open()
	if err != nil {
		return w, h
	}
	defer ttyIn.Close()
	osW, osH, err := ttyIn.Size()
	if err != nil {
		return w, h
	}
	osW, osH = osW-1, osH-1

	// 从起点位置开始算， 计算能放多少格， 以第一格为准
	if t.opt.AutoWidth {
		w = 0
		for col := org.Col; col < end.Col; col++ {
			if col >= len(t.body[0]) {
				break
			}
			// 这里-4是为了让表格线有空间写入
			if osW -= t.body[0][col].ColWidth() + 4; osW >= 0 {
				w += 1
			}
		}
	}
	if t.opt.AutoHeight {
		h = 0
		for row := org.Row; row < end.Row; row++ {
			if row >= len(t.body) || len(t.body[row]) == 0 {
				break
			}
			// 这里-2是为了让表格线有空间
			if osH -= t.body[row][0].RowHeight() + 2; osH >= 0 {
				h += 1
			}
		}
	}

	return w, h
}
