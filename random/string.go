package random

import (
	"crypto/rand"
	"fmt"
	rand2 "math/rand"
)

// GetRandString get random string use crypto/rand
func GetRandString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return GetRandStringNormal(n)
	}
	return fmt.Sprintf("%x", b)[:n]
}

// GetRandStringNormal get random string use own seeds
func GetRandStringNormal(n int) string {
	b := make([]byte, n)
	var runes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := range b {
		b[i] = runes[rand2.Intn(len(runes))]
	}
	return fmt.Sprintf("%x", b)[:n]
}
