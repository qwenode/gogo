package ss

import (
	"strings"
)

// Substr with length
func Substr(str string, start, length int) string {
	if length < 0 {
		length = 0
	}
	s := []rune(str)
	sLen := len(s)
	begin := start
	if start < 0 {
		begin = sLen + start
	}
	if begin < 0 {
		begin = 0
	}

	end := begin + length
	if end > sLen {
		end = sLen
	}
	if begin == 0 && end == sLen {
		return str
	}
	return string(s[begin:end])
}

// CutByEndString input str:"Pure organic linen fabrics, clothes and roses that last by LinenRoses" endStr:"by"
// return: "Pure organic linen fabrics, clothes and roses that last"
func CutByEndString(str, endStr string) string {
	split := strings.Split(str, endStr)
	end := len(split) - 1
	if end < 1 {
		end = 1
	}
	return strings.TrimSpace(strings.Join(split[:end], endStr))
}

// GetLastElemBySep input str: "ab,bc,cc", sep:"," return :"cc"
func GetLastElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) <= 1 {
		return str
	}
	return strings.TrimSpace(split[len(split)-1])
}

// GetFirstElemBySep input str: "ab,bc,cc", sep:"," return :"ab"
func GetFirstElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) == 1 {
		return str
	}
	return strings.TrimSpace(split[0])
}

// GetSecondElemBySep input str: "ab,bc,cc", sep:"," return :"bc"
func GetSecondElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) < 2 {
		return str
	}
	return strings.TrimSpace(split[1])
}

// JoinWithSep merge string with separator
func JoinWithSep(sep string, elem ...string) string {
	if len(elem) == 0 {
		return ""
	}
	return strings.Trim(strings.Join(elem, sep), sep)
}

// Contains find any string in the str
func Contains(str string, finds ...string) bool {
	for _, find := range finds {
		if strings.Contains(str, find) {
			return true
		}
	}
	return false
}

// ContainsAllOfAll matches any things
func ContainsAllOfAll(str string, finds ...string) bool {
	str = strings.ToLower(str)
	for i, find := range finds {
		finds[i] = strings.ToLower(find)
	}
	for _, find := range finds {
		if strings.Contains(str, find) || strings.Contains(find, str) {
			return true
		}
	}
	return false
}

// Join merge string
func Join(elem ...string) string {
	if len(elem) == 0 {
		return ""
	}
	return strings.Join(elem, "")
}
