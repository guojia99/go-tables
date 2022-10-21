package utils

import "math"

func Max(in ...int) int {
	if len(in) == 0 {
		return -1
	}
	cache := -math.MaxInt
	for _, val := range in {
		if val > cache {
			cache = val
		}
	}
	return cache
}

func UintMax(in ...uint) uint {
	cache := uint(0)
	for _, val := range in {
		if val > cache {
			cache = val
		}
	}
	return cache
}
