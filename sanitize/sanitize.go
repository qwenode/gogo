package sanitize

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func Int(str string) int {
	sanitize, _ := regexp.Compile("([0-9]+)")
	r := sanitize.FindAllString(str, -1)
	i, _ := strconv.Atoi(strings.Join(r, ""))
	return i
}

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

func Float64(str string) float64 {
	sanitize, _ := regexp.Compile("([0-9.]+)")
	r := sanitize.FindAllString(str, -1)
	float, _ := strconv.ParseFloat(strings.Join(r, ""), 64)
	return float
}

func Alphabet(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z ]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func AlphabetWithoutSpace(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	return strings.TrimSpace(s)
}

func AlphabetNumber(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9 ]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	s = strings.Join(strings.Fields(s), " ")
	return s
}

func AlphabetNumberWithoutSpace(str string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9]+)")
	r := sanitize.FindAllString(str, -1)
	s := strings.Join(r, "")
	return strings.TrimSpace(s)
}

// ATTENTION! if fail default return "/"
func DirectoryPath(path string) string {
	sanitize, _ := regexp.Compile("([a-zA-Z0-9 \\-_]+)")
	r := sanitize.FindAllString(path, -1)
	s := strings.Join(r, "/")
	return "/" + strings.TrimLeft(strings.TrimSpace(s), "/")
}
