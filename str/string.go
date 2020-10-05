package str

func Substr(str string, start, length int) string {
	if start < 0 {
		start = 0
	}
	s := []rune(str)
	sLen := len(s)
	end := start + length
	if end > sLen {
		return str
	}
	if start == 0 && end == sLen {
		return str
	}
	return string(s[start:end])
}
