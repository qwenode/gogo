package hashes

import (
	md52 "crypto/md5"
	"fmt"
)

// Md5 get string md5 hash
func Md5(str string) string {
	return Md5Byte([]byte(str))
}

// Md5Byte get byte md5 hash
func Md5Byte(data []byte) string {
	return fmt.Sprintf("%x", md52.Sum(data))
}
