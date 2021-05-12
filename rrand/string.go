package rrand

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
	"time"
)

// GetRandString get rrand string use crypto/rand
func GetRandString(n int) string {
	b := make([]byte, n)
	if _, err := crypto_rand.Read(b); err != nil {
		return GetRandStringNormal(n)
	}
	return fmt.Sprintf("%x", b)[:n]
}

// GetRandStringNormal get rrand string use own seeds
func GetRandStringNormal(n int) string {
	b := make([]byte, n)
	var runes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	for i := range b {
		b[i] = runes[math_rand.Intn(len(runes))]
	}
	return fmt.Sprintf("%x", b)[:n]
}

func init() {
	var rb [4]byte
	_, err := crypto_rand.Read(rb[:])
	seed := time.Now().UnixNano()
	if err == nil {
		seed += int64(binary.LittleEndian.Uint32(rb[:]))
	}
	math_rand.Seed(seed)
}
