package table

import (
	"fmt"
	"github.com/guojia99/go-tables/_example/zen"
	"testing"
)

func TestAlign_Repeat(t *testing.T) {
	a := AlignTopLeft

	fmt.Println(zen.List[0])
	fmt.Println(a.Repeat(zen.List[0], 10))
}
