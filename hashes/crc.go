package hashes

import (
	"fmt"
	"hash/crc32"
)

func Crc32(str string) string {
	table := crc32.MakeTable(crc32.IEEE)
	byteString := crc32.Checksum([]byte(str), table)
	return fmt.Sprintf("%x", byteString)
}
