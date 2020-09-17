package hashes

import (
	sha12 "crypto/sha1"
	"fmt"
)

func Sha1(str string) string {
	hash := sha12.New()
	hash.Write([]byte(str))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)
}
