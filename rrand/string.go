package rrand

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
	"strconv"
	"strings"
	"time"
)

// GetRandString get rrand string use crypto/rand
func GetRandString(n int) string {
	b := make([]byte, n)
	_, _ = crypto_rand.Read(b)
	return fmt.Sprintf("%x", b)[:n]
}

// GetRandStringNormal get rrand string use own seeds
func GetRandStringNormal(n int) string {
	b := ""
	var runes = strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	for i := 0; i < n; i++ {
		b += runes[math_rand.Intn(len(runes))]
	}
	return b
}
func GetRandNumber(length int) string {
	b := ""
	runes := Perm(10)
	for i := 0; i < length; i++ {
		r := math_rand.Intn(len(runes))
		b += strconv.Itoa(runes[r])
	}
	return b
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
