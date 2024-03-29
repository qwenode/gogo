package uurl

import (
	"net/url"
	"strings"
)

// UrlDecode url.QueryUnescape
// Deprecated
func UrlDecode(s string) string {
	unescape, _ := url.QueryUnescape(s)
	return unescape
}

// IsUrl check url
// Deprecated
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
