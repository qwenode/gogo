package hashes

import (
	sha12 "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha1(str string) string {
	hash := sha12.New()
	hash.Write([]byte(str))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)
}

func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)
}
func Sha512(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	byteString := hash.Sum(nil)
	return fmt.Sprintf("%x", byteString)
}
