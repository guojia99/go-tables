package table

import (
	"fmt"
	"testing"
)

func TestTable_String(t1 *testing.T) {
	//data := [][]int{
	//	{112, 2, 3},
	//	{1, 223, 3},
	//	{4, 5, 6232132},
	//	{72312312, 2328, 921312},
	//}
	data := []struct {
		A int
		B int
		c int
	}{
		{112, 2, 3},
		{1, 223, 3},
		{4, 5, 6232132},
		{72312312, 2328, 921312},
	}

	opt1 := &Option{
		Align:   AlignCenter,
		Contour: DefaultContour,
	}

	opt2 := &Option{
		Align:   AlignCenter,
		Contour: MariaDBContour,
	}

	opt3 := &Option{
		Align:   AlignCenter,
		Contour: EmptyContour,
	}
	tb, _ := SimpleTable(data, opt1)
	fmt.Println(tb)

	tb2, _ := SimpleTable(data, opt2)
	fmt.Println(tb2)

	tb3, _ := SimpleTable(data, opt3)
	fmt.Println(tb3)
}
