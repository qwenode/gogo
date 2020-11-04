package hashes

import (
	md52 "crypto/md5"
	"fmt"
)

// Md5 get string md5 hash
func Md5(str string) string {
	return fmt.Sprintf("%x", md52.Sum([]byte(str)))
}
