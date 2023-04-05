package uurl

import "net/url"

// UrlDecode url.QueryUnescape
func UrlDecode(s string) string {
    unescape, _ := url.QueryUnescape(s)
    return unescape
}

// IsUrl check url
func IsUrl(s string) bool {
    _, err := url.Parse(s)
    if err != nil {
        return false
    }
    return true
}
