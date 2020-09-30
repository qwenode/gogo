package hashes

import (
	md52 "crypto/md5"
	"fmt"
)

func Md5(str string) string {
	byteStr := md52.Sum([]byte(str))
	return fmt.Sprintf("%x", byteStr)
}
