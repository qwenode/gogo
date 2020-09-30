package hashes

import (
	sha12 "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha1(str string) string {
	byteString := sha12.Sum([]byte(str))
	return fmt.Sprintf("%x", byteString)
}

func Sha256(str string) string {
	byteString := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", byteString)
}
func Sha512(str string) string {
	byteString := sha512.Sum512([]byte(str))
	return fmt.Sprintf("%x", byteString)
}
