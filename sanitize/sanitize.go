package sanitize

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// Int sanitize string to int ,exp: 1s2d3 out:123
func Int(str string) int {
	sanitize, _ := regexp.Compile("([0-9]+)")
	r := sanitize.FindAllString(str, -1)
	i, _ := strconv.Atoi(strings.Join(r, ""))
	if strings.Index(str, "-") == 0 && i > 0 {
		i *= -1
	}
	return i
}

// UInt sanitize string to UInt ,exp: 1s2d3 out:123
func UInt(str string) uint {
	sanitize, _ := regexp.Compile("([0-9]+)")
	r := sanitize.FindAllString(str, -1)
	i, _ := strconv.Atoi(strings.Join(r, ""))
	return uint(i)
}

// HostName get hostname by given string
func HostName(u string) string {
	u = strings.TrimSpace(u)
	if strings.Index(u, "://") < 0 {
		u = "http://" + u
	}
	parse, _ := url.Parse(u)
	if parse == nil {
		return ""
	}
	host := parse.Hostname()
	isWide := strings.Index(host, "*.") == 0
	match := regexp.MustCompile(`([a-zA-Z0-9.-]+)`)
	r := strings.Trim(match.FindString(host), ".")
	if r == "" {
		return ""
	}
	if isWide {
		r = "*." + r
	}
	return r
}

// Float64 sanitize string to float
func Float64(str string) float64 {
	sanitize, _ := regexp.Compile("([0-9.]+)")
	r := sanitize.FindAllString(str, -1)
	float, _ := strconv.ParseFloat(strings.Join(r, ""), 64)
	return float
}

// Alphabet sanitize string , exp: 1s2s3 out:ss
func Alphabet(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z ]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

// AlphabetWithoutSpace sanitize string , exp: 1s2s 3# out:ss
func AlphabetWithoutSpace(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	return strings.TrimSpace(s)
}

// AlphabetNumber sanitize string , exp: 1s2s 3# out:1s2s 3
func AlphabetNumber(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9 ]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

// AlphabetNumberWithoutSpace sanitize string , exp: 1s2s 3# out:1s2s3
func AlphabetNumberWithoutSpace(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	return strings.TrimSpace(s)
}

// DirectoryPath ATTENTION! if fail default return "/"
func DirectoryPath(path string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9\\-_]+)")
	r := sanitize.FindAllString(path, -1)
	s := strings.Join(r, "/")
	return "/" + strings.TrimLeft(strings.TrimSpace(s), "/")
}

// FilePath ATTENTION! if fail default return "/"
func FilePath(path string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9\\-_\\.]+)")
	r := sanitize.FindAllString(path, -1)
	s := strings.Join(r, "/")
	return "/" + strings.TrimLeft(strings.TrimSpace(s), "/")
}

// MultipleSpaceToSingle aa  bbb to aa bbb
func MultipleSpaceToSingle(str string) string {
	sanitize, _ := regexp.Compile("\\s+")
	r := sanitize.ReplaceAllString(str, " ")
	return strings.TrimSpace(r)
}

//MultipleBackslashToSingle \\ to \  \\\\ to \
func MultipleBackslashToSingle(str string) string {
	sanitize, _ := regexp.Compile("([\\\\]+)")
	r := sanitize.ReplaceAllString(str, "\\")
	return strings.TrimSpace(r)
}

// RemoveAllSpace replace aa, bb to aa,bb
func RemoveAllSpace(str string) string {
	sanitize, _ := regexp.Compile("\\s+")
	r := sanitize.ReplaceAllString(str, "")
	return strings.TrimSpace(r)
}

// StripHtml remove > < %
func StripHtml(str string) string {
	str = strings.ReplaceAll(str, "<", "")
	str = strings.ReplaceAll(str, ">", "")
	str = strings.ReplaceAll(str, "%3e", "")
	str = strings.ReplaceAll(str, "%3E", "")
	str = strings.ReplaceAll(str, "%3C", "")
	str = strings.ReplaceAll(str, "%3c", "")
	return strings.TrimSpace(str)
}
