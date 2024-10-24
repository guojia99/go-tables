package tables

import (
	"fmt"
	"testing"
)

func _testTable() *table {
	return &table{
		opt: Option{
			Contour: DefaultContour,
		},
		body: Cells2D{
			{NewCell(1), NewCell(2), NewCell(3)},
			{NewCell(4), NewCell(5), NewCell(6)},
			{NewCell(7), NewCell(8), NewCell(9)},
		},
	}
}

func Test_table_DeleteRows(t1 *testing.T) {
	t1.Run(
		"delete has", func(t *testing.T) {
			tb := _testTable()
			if _, err := tb.DeleteRow(2); err != nil {
				t1.Fatal(err)
			}
			fmt.Println(tb.body)
		},
	)
}

func Test_table_DeleteCol(t1 *testing.T) {
	t1.Run(
		"delete has", func(t *testing.T) {
			tb := _testTable()
			if _, err := tb.DeleteCol(1); err != nil {
				t1.Fatal(err)
			}
			fmt.Println(tb.body)
		},
	)
}

func Test_table_String(t1 *testing.T) {
	//tb := _testTable()
	//fmt.Println(tb.String())

	tb2 := NewTable()
	tb2.SetOption(
		Option{
			Contour: DefaultContour,
		},
	)
	tb2.AddBody(NewCell("123", "321"), NewCell("33122", "2131231", "2131231"))
	tb2.AddBody(NewCell("123", "321"), NewCell("33122", "2131231", "2131231"))

	_ = tb2.SetColWidth(10)
	_ = tb2.SetRowHeight(3)
	_ = tb2.SetCellAlignByRow(AlignCenter)

	fmt.Println(tb2.String())
}
