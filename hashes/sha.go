package hashes

import (
	sha12 "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Sha1 get string sha1 hash
func Sha1(str string) string {
	return fmt.Sprintf("%x", sha12.Sum([]byte(str)))
}

// Sha256 get string sha256 hash
func Sha256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}

// Sha512 get string sha512 hash
func Sha512(str string) string {
	return fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
}
