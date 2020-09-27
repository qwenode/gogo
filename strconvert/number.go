package strconvert

import "strconv"

func ToInt(str string) int {
	atoi, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return atoi
}
