package tables

import (
	"fmt"
	"testing"

	zen "github.com/guojia99/go-tables/zen"
)

func Test_Cell_String(t *testing.T) {
	fmt.Println(NewCell(zen.ZenList, AlignCenter))
	fmt.Println(NewCell(zen.ZenListZn, AlignCenter))
}
