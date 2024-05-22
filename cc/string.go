package cc

import (
	"fmt"
	"strconv"
)

// ToString convert to string
func ToString(str interface{}) string {
	switch str.(type) {
	case int:
		return strconv.Itoa(str.(int))
	case int32:
		return fmt.Sprintf("%v", str.(int32))
	case int64:
		return fmt.Sprintf("%v", str.(int64))
	case string:
		return str.(string)
	case float64:
		return fmt.Sprintf("%v", str.(float64))
	case float32:
		return fmt.Sprintf("%v", str.(float32))
	default:
		return ""
	}
}
