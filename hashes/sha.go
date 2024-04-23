package hashes

import (
	sha12 "crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Sha1 get string sha1 hash
func Sha1(str string) string {
	return Sha1Byte([]byte(str))
}

// Sha1Byte get byte sha1 hash
func Sha1Byte(data []byte) string {
	return fmt.Sprintf("%x", sha12.Sum(data))
}

// Sha256 get string sha256 hash
func Sha256(str string) string {
	return Sha256Byte([]byte(str))
}

// Sha256Byte get bytes sha256 hash
func Sha256Byte(data []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(data))
}

// Sha512 get string sha512 hash
func Sha512(str string) string {
	return Sha1Byte([]byte(str))
}

// Sha512Byte get bytes sha512 hash
func Sha512Byte(data []byte) string {
	return fmt.Sprintf("%x", sha512.Sum512(data))
}
