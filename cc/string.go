package cc

import (
    "fmt"
    "strconv"
    "strings"
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

// ToBoolean convert to boolean
func ToBoolean(data interface{}) bool {
    switch data.(type) {
    case int, int32, int64, uint, uint32, uint64, uint8, int8:
        toInt := ToInt(data)
        return toInt > 0
    case string:
        s := strings.TrimSpace(strings.ToLower(data.(string)))
        if s == "1" || s == "ok" || s == "true" || s == "y" || s == "yes" || s == "t" {
            return true
        }
        return false
    case bool:
        return data.(bool)
    }
    return false
}
