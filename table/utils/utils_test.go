package utils

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	i := Max(1, 2, 4, -43, 55)
	fmt.Println(i)
}
