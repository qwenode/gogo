package hashes

import (
	md52 "crypto/md5"
	"fmt"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md52.Sum([]byte(str)))
}
