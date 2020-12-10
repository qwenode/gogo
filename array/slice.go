package array

import "sort"

// IntInSlice Check int Element Exists in a Slice
func IntInSlice(search int, slice []int) bool {
	for _, i := range slice {
		if i == search {
			return true
		}
	}
	return false
}

// StringInSlice Check int Element Exists in a Slice
func StringInSlice(search string, slice []string) bool {
	for _, i := range slice {
		if i == search {
			return true
		}
	}
	return false
}

// IntegerRange use for MergeIntRange
type IntegerRange struct {
	Low  int
	High int
}

// MergeIntRange input [1,3], [2,6], [5,10] [2] becomes [1,10]
// input [11,13], [2,6], [5,10] [2,0] [30,0] becomes [1,10] [11,13] [30,0]
func MergeIntRange(data []IntegerRange) []IntegerRange {
	m := append([]IntegerRange(nil), data...)
	if len(m) <= 1 {
		return m
	}

	sort.Slice(m,
		func(i, j int) bool {
			if m[i].Low < m[j].Low {
				return true
			}
			if m[i].Low <= m[j].Low && m[i].High > m[j].High {
				return true
			}
			if m[i].Low == m[j].Low && m[i].High <= m[j].High {
				return true
			}
			return false
		},
	)

	j := 0
	for i := 1; i < len(m); i++ {
		if m[j].High >= m[i].Low {
			if m[j].High < m[i].High {
				m[j].High = m[i].High
			}
		} else if m[j].High == m[i].High && m[j].Low == m[i].Low {
			m[j] = m[i]
		} else {
			j++
			m[j] = m[i]
		}

	}

	return append([]IntegerRange(nil), m[:j+1]...)
}
