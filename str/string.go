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

func CutByEndString(str, endStr string) string {
	split := strings.Split(str, endStr)
	end := len(split) - 1
	if end < 1 {
		end = 1
	}
	return strings.TrimSpace(strings.Join(split[:end], endStr))
}
