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

func (t *table) getWidthHeightAndOrgPoint() (int, int, Address) {
	// 重新计算长度
	org := t.opt.OrgPoint
	end := t.opt.EndPoint

	if end.Row == end.Col && end.Row == 0 {
		end.Col, end.Row = len(t.body[0]), len(t.body)
	}

	if end.Col == org.Col || end.Col == 0 || end.Col-org.Col <= 0 {
		org.Col, end.Col = 0, len(t.body[0])
	}
	if end.Row == org.Row || end.Row == 0 || end.Row-org.Row <= 0 {
		org.Row, end.Row = 0, len(t.body)
	}

	// 正常情况
	w := end.Col - org.Col
	h := end.Row - org.Row
	if !t.opt.AutoWidth && !t.opt.AutoHeight {
		return w, h, org
	}

	// 依据屏幕的情况
	ttyIn, err := tty.Open()
	if err != nil {
		return w, h, org
	}
	defer ttyIn.Close()
	osW, osH, err := ttyIn.Size()
	if err != nil {
		return w, h, org
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

	return w, h, org
}

// todo 没写完
func (t *table) parserTableContent() (out []string) {
	t.autoScaling()
	// 同步最大行数

	var width, height, org = t.getWidthHeightAndOrgPoint() // 这是格子数
	/*
		             ↓
			[x] [x] [x] [x] [x] [x] [x] [x]
			[x] [x] [x] [x] [x] [x] [x] [x]
		 →	[x] [x] [o] [1] [1] [1] [1] [x]
			[x] [x] [1] [1] [1] [1] [1] [x]
			[x] [x] [1] [1] [1] [1] [e] [x] ←
			[x] [x] [x] [x] [x] [x] [x] [x]
									 ↑
			将从原点开始获取
			[o] [1] [1] [1] [1]
			[1] [1] [1] [1] [1]
			[1] [1] [1] [1] [e]

		 	获取出来将是
		    t.opt.EndPoint - t.opt.OrgPoint
	*/
	var colWidths = make([]int, 0)
	for row := org.Row; row < org.Row+height; row++ {
		// 校验body
		if row >= len(t.body) || len(t.body[row]) == 0 {
			break
		}

		// 头一行 负责计算宽度
		if row == org.Row {
			for col := org.Col; col < org.Col+width; col++ {
				cell := t.body[row][col]
				colWidths = append(colWidths, cell.ColWidth())
				// 把这一行的数据塞进去
			}
			out = append(out, t.opt.Contour.Handler(colWidths))
		}

		// 填充行数据
		var cell0 = t.body[row][org.Col]

		var msgs = make([][]string, cell0.RowHeight())
		for i := 0; i < cell0.RowHeight(); i++ {
			msgs[i] = make([]string, width-org.Col)
		}
		for col := org.Col; col < org.Col+width; col++ {
			lines := t.body[row][col].Align().Repeats(t.body[row][col].Lines(), t.body[row][col].ColWidth())
			for idx, line := range lines {
				if idx > len(msgs) {
					continue
				}
				msgs[idx][col] = line
			}
		}

		for _, msg := range msgs {
			out = append(out, t.opt.Contour.Content(msg))
		}

		// 首行作为头, 需要加入一份断行
		if row == org.Row {
			out = append(out, t.opt.Contour.Intersection(colWidths))
		}
	}

	out = append(out, t.opt.Contour.Footer(colWidths))
	return out
}

func (t *table) getColMaxWidth(col int) int {
	var maxCol = 0

	for row := range t.body {
		if col >= len(t.body[row]) {
			continue
		}

		if maxCol < t.body[row][col].ColWidth() {
			maxCol = t.body[row][col].ColWidth()
		}
	}
	return maxCol
}
