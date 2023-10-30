package ss

import (
	"net/url"
	"regexp"
	"strings"
)

// UrlDecode url.QueryUnescape
func UrlDecode(s string) string {
	unescape, _ := url.QueryUnescape(s)
	return unescape
}

// IsUrl check url
func IsUrl(s string) bool {
	if strings.TrimSpace(s) == "" {
		return false
	}
	_, err := url.Parse(s)
	if err != nil {
		return false
	}
	return true
}
func ExtractUrls(str string) []string {
	if len(str) <= 0 {
		return []string{}
	}
	compile, _ := regexp.Compile(`\bhttps?://[^,\s()<>]+(?:\([\w\d]+\)|([^,[:punct:]\s]|/))`)
	submatch := compile.FindAllString(str, -1)
	return submatch
}
