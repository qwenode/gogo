package mmath

import "math"

// Round
func Round(f float64, n int) float64 {
	pow10N := math.Pow10(n)
	return math.Trunc((f+0.5/pow10N)*pow10N) / pow10N
}

// CeilInt ceil int for pagination
func CeilInt(first, second int) int {
	if first <= 0 {
		return 0
	}
	if second <= 0 {
		return first
	}
	return int(math.Ceil(float64(first) / float64(second)))
}

// CeilInt64 ceil int64 for pagination
func CeilInt64(first, second int64) int64 {
	if first <= 0 {
		return 0
	}
	if second <= 0 {
		return first
	}
	return int64(math.Ceil(float64(first) / float64(second)))
}

// BetweenInt return (val >= start && val <= end)
func BetweenInt(val, start, end int) bool {
	return val >= start && val <= end
}

// BetweenInt64 val >= start && val <= end
func BetweenInt64(val, start, end int64) bool {
	return val >= start && val <= end
}
