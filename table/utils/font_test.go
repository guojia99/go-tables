package utils

import (
	"fmt"
	"testing"
)

func TestRealLength(t *testing.T) {
	b := []byte{0, 66, 68, 70, 0, 0, 50, 52, 54, 52, 55, 49, 0, 0, 0, 71, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println(RealLength(string(b)))

	in := string(b)
	fmt.Println(RealLength(in))
}
