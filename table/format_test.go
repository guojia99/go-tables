package table

import (
	"fmt"
	"testing"
)

func TestContour_Handler(t *testing.T) {
	sw := []int{4, 4, 5, 1, 2, 5}
	fmt.Println(DefaultContour.Handler(sw))
	fmt.Println(DefaultContour.Footer(sw))
}
