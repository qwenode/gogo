package str

import "strings"

func Substr(str string, start, length int) string {

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

// input str:"Pure organic linen fabrics, clothes and roses that last by LinenRoses" endStr:"by"
// return: "Pure organic linen fabrics, clothes and roses that last"
func CutByEndString(str, endStr string) string {
	split := strings.Split(str, endStr)
	end := len(split) - 1
	if end < 1 {
		end = 1
	}
	return strings.TrimSpace(strings.Join(split[:end], endStr))
}

// input str: "ab,bc,cc", sep:"," return :"cc"
func GetLastElemBySep(str, sep string) string {
	split := strings.Split(str, sep)
	if len(split) <= 1 {
		return str
	}
	return strings.TrimSpace(split[len(split)-1])
}
