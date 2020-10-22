package convert

import "strconv"

func ToInt(str interface{}) int {
	switch str.(type) {
	case uint:
		return int(str.(uint))
	case uint64:
		return int(str.(uint64))
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

func ToUInt(str interface{}) uint {
	switch str.(type) {
	case uint64:
		return uint(str.(uint64))
	case uint:
		return str.(uint)
	case int:
		return uint(str.(int))
	case int64:
		return uint(str.(int64))
	case string:
		atoi, _ := strconv.Atoi(str.(string))
		return uint(atoi)
	case float64:
		return uint(str.(float64))
	case float32:
		return uint(str.(float32))
	default:
		return 0
	}
}
func ToUInt64(str interface{}) uint64 {
	switch str.(type) {
	case uint64:
		return str.(uint64)
	case uint:
		return uint64(str.(uint))
	case int:
		return uint64(str.(int))
	case int64:
		return uint64(str.(int64))
	case string:
		atoi, _ := strconv.Atoi(str.(string))
		return uint64(atoi)
	case float64:
		return uint64(str.(float64))
	case float32:
		return uint64(str.(float32))
	default:
		return 0
	}
}
func ToInt64(str interface{}) int64 {
	switch str.(type) {
	case uint64:
		return int64(str.(uint64))
	case uint:
		return int64(str.(uint))
	case int:
		return int64(str.(int))
	case int64:
		return str.(int64)
	case string:
		atoi, _ := strconv.Atoi(str.(string))
		return int64(atoi)
	case float64:
		return int64(str.(float64))
	case float32:
		return int64(str.(float32))
	default:
		return 0
	}
}

func ToFloat64(str interface{}) float64 {
	switch str.(type) {
	case uint64:
		return float64(str.(uint64))
	case uint:
		return float64(str.(uint))
	case int:
		return float64(str.(int))
	case int64:
		return float64(str.(int64))
	case string:
		float, _ := strconv.ParseFloat(ToString(str), 64)
		return float
	case float64:
		return str.(float64)
	case float32:
		return float64(str.(float32))
	default:
		return float64(0)
	}
}
