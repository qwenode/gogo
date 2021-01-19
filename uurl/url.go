package uurl

import "net/url"

// UrlDecode url.QueryUnescape
func UrlDecode(s string) string {
	unescape, _ := url.QueryUnescape(s)
	return unescape
}
