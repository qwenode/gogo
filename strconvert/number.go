package strconvert

import "strconv"

func ToInt(str interface{}) int {
	switch str.(type) {
	case int:
		return str.(int)
	case int64:
		return int(str.(int64))
	case string:
		atoi, _ := strconv.Atoi(str.(string))
		return atoi
	case float64:
		return int(str.(float64))
	case float32:
		return int(str.(float32))
	default:
		return 0
	}
}
