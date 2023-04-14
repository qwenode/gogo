package mmap

// GetMapIntKeys get map keys,return slice
// Deprecated see GetIntegerKeys
func GetMapIntKeys(m map[int]interface{}) []int {
    return GetIntegerKeys(m)
}

// GetIntegerKeys get map keys,return slice
func GetIntegerKeys(m map[int]interface{}) []int {
    keys := make([]int, 0, len(m))
    for i := range m {
        keys = append(keys, i)
    }
    return keys
}

// GetStringKeys get map keys,return slice
func GetStringKeys(m map[string]interface{}) []string {
    keys := make([]string, 0, len(m))
    for i := range m {
        keys = append(keys, i)
    }
    return keys
}
