package sanitize

import (
	"regexp"
	"strconv"
	"strings"
)

func Int(str string) int {
	sanitize, _ := regexp.Compile("([0-9]+)")
	r := sanitize.FindAllString(str, -1)
	//log.Println(r)
	i, _ := strconv.Atoi(strings.Join(r, ""))
	return i
}
func Float64(str string) float64 {
	sanitize, _ := regexp.Compile("([0-9.]+)")
	r := sanitize.FindAllString(str, -1)
	//log.Println(r)
	float, _ := strconv.ParseFloat(strings.Join(r, ""), 64)
	return float
}
