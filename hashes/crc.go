package hashes

import (
	"fmt"
	"hash/crc32"
)

// Crc32 get string crc32 hash
func Crc32(str string) string {
	return Crc32Byte([]byte(str))
}

// Crc32Byte get byte crc32 hash
func Crc32Byte(data []byte) string {
	table := crc32.MakeTable(crc32.IEEE)
	byteString := crc32.Checksum(data, table)
	return fmt.Sprintf("%x", byteString)
}
