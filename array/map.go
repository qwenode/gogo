package array

// GetMapIntKeys get map keys,return slice
func GetMapIntKeys(m map[int]interface{}) []int {
	keys := make([]int, 0, len(m))
	for i := range m {
		keys = append(keys, i)
	}
	return keys
}
