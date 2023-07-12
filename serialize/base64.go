package serialize

import "encoding/base64"

// Base64DecodeByte decode base64 string to byte
func Base64DecodeByte(s string) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(s)
	return decodeString
}

// Base64Decode decode base64 string
func Base64Decode(s string) string {
	return string(Base64DecodeByte(s))
}
