package array

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
