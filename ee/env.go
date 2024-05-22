package ee

import (
	"github.com/qwenode/gogo/cc"
	"os"
	"strings"
)

func GetBool(key string) bool {
	value := strings.ToLower(GetValue(key))
	switch value {
	case "ok":
	case "true":
	case "1":
	case "yes":
	case "y":
	case "on":
		return true
	}
	return false
}
func GetStringList(key string) []string {
	value := GetValue(key)
	split := strings.Split(value, ",")
	list := make([]string, 0)
	for _, s := range split {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		list = append(list, s)
	}
	return list
}
func GetIntList(key string) []int {
	value := GetValue(key)
	split := strings.Split(value, ",")
	list := make([]int, 0)
	for _, s := range split {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		toInt := cc.ToInt(s)
		list = append(list, toInt)
	}
	return list
}
func GetInt(key string) int {
	return cc.ToInt(GetValue(key))
}
func GetInt32(key string) int32 {
	return cc.ToInt32(GetValue(key))
}
func GetFloat64(key string) float64 {
	return cc.ToFloat64(GetValue(key))
}
func GetString(key string) string {
	return GetValue(key)
}
func GetInt64(key string) int64 {
	return cc.ToInt64(GetValue(key))
}
func GetValue(key string) string {
	getenv := os.Getenv(key)
	return strings.TrimSpace(getenv)
}
