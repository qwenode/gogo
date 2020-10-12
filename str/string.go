package str

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
