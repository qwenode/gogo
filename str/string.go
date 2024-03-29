package str

import (
	"regexp"
	"strings"
)

// Deprecated
func ExtractUrls(str string) []string {
	if len(str) <= 0 {
		return []string{}
	}
	compile, _ := regexp.Compile(`\bhttps?://[^,\s()<>]+(?:\([\w\d]+\)|([^,[:punct:]\s]|/))`)
	submatch := compile.FindAllString(str, -1)
	return submatch
}

// Deprecated with length
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
// Deprecated
func CutByEndString(str, endStr string) string {
	split := strings.Split(str, endStr)
	end := len(split) - 1
	if end < 1 {
		end = 1
	}
	return strings.TrimSpace(strings.Join(split[:end], endStr))
}

// GetLastElemBySep input str: "ab,bc,cc", sep:"," return :"cc"
//Deprecated
func GetLastElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) <= 1 {
		return str
	}
	return strings.TrimSpace(split[len(split)-1])
}

// GetFirstElemBySep input str: "ab,bc,cc", sep:"," return :"ab"
//Deprecated
func GetFirstElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) == 1 {
		return str
	}
	return strings.TrimSpace(split[0])
}

// GetSecondElemBySep input str: "ab,bc,cc", sep:"," return :"bc"
//Deprecated
func GetSecondElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) < 2 {
		return str
	}
	return strings.TrimSpace(split[1])
}

// JoinWithSep merge string with separator
//Deprecated
func JoinWithSep(sep string, elem ...string) string {
	if len(elem) == 0 {
		return ""
	}
	return strings.Trim(strings.Join(elem, sep), sep)
}

// Join merge string
//Deprecated
func Join(elem ...string) string {
	if len(elem) == 0 {
		return ""
	}
	return strings.Join(elem, "")
}
