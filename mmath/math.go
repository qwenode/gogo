package mmath

import "math"

// ceil int for pagination
func CeilInt(first, second int) int {
	if first <= 0 {
		return 0
	}
	if second <= 0 {
		return first
	}
	return int(math.Ceil(float64(first) / float64(second)))
}

// ceil int64 for pagination
func CeilInt64(first, second int64) int64 {
	if first <= 0 {
		return 0
	}
	if second <= 0 {
		return first
	}
	return int64(math.Ceil(float64(first) / float64(second)))
}
