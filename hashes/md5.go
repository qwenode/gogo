package hashes

import (
	md52 "crypto/md5"
	"fmt"
)

func Md5(str string) string {
	hash := md52.New()
	hash.Write([]byte(str))
	byteStr := hash.Sum(nil)
	return fmt.Sprintf("%x", byteStr)
}
