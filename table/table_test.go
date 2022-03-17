package table

import (
	"testing"
)

func TestTable_String(t1 *testing.T) {
	data := [][]int{
		{112, 2, 3},
		{1, 223, 3},
		{4, 5, 6232132},
		{72312312, 2328, 921312},
	}

	tb, err := SimpleTable(data, &Option{
		Align: AlignCenter,
	})

	if err != nil {
		t1.Errorf("%v", err)
	}
	_ = tb.String()
}
