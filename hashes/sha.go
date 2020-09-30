package hashes

import (
	sha12 "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha12.Sum([]byte(str)))
}

func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}
