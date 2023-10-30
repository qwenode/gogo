package file

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"github.com/qwenode/gogo/ss"
	"github.com/qwenode/gogo/str"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// GetFileNameWithoutExtension get file name without extension exp: xxx/aa.go returns: aa
// Deprecated
func GetFileNameWithoutExtension(filePath string) string {
	filePath = strings.TrimSpace(filePath)
	if len(filePath) == 0 {
		return filePath
	}
	base := filepath.Base(filePath)
	return ss.GetFirstElemBySep(base, ".")
}

// GetExtension get extension name with no check mime type
// Deprecated
func GetExtension(filePath string) string {
	return str.GetLastElemBySep(filePath, ".")
}

// GetContents get file content
// Deprecated
func GetContents(filename string) string {
	return string(GetContentsByte(filename))
}

// GetContentsByte get file content returns byte
// Deprecated
func GetContentsByte(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}
	}
	return file
}

// GetLines get every line of file
// Deprecated
func GetLines(filename string) (lines []string) {
	return strings.Split(GetContents(filename), "\n")
}

// Exist check file or dir if exists
// Deprecated
func Exist(filename string) bool {
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Size get file size
// Deprecated
func Size(filename string) int64 {
	f, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return f.Size()
}

// WriteFileAppend Append content to the end of the file
// Deprecated
func WriteFileAppend(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err1 := f.Close(); err1 == nil {
		err = err1
	}

	return err
}

// PutContents put content to the file,will clean old data
// Deprecated
func PutContents(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err1 := f.Close(); err1 == nil {
		err = err1
	}

	return err
}

// Sha1 get file sha1 hash
// Deprecated
func Sha1(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Sha256 get file sha256 hash
// Deprecated
func Sha256(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Md5 get file md5 hash
// Deprecated
func Md5(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Crc32 get file crc32 hash
// Deprecated
func Crc32(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	table := crc32.MakeTable(crc32.IEEE)
	hash32 := crc32.New(table)
	if _, err := io.Copy(hash32, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash32.Sum(nil)), nil
}

// GetSize get file size
// Deprecated
func GetSize(fileName string) int64 {
	stat, err := os.Stat(fileName)
	if err != nil {
		return 0
	}
	return stat.Size()
}

// IsDirectory check file is directory
// Deprecated
func IsDirectory(filename string) bool {
	stat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// CopyFile copy file
// Deprecated
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

// MoveFile move file
// Deprecated
func MoveFile(src, dst string) error {
	err := CopyFile(src, dst)
	if err != nil {
		return err
	}
	return os.Remove(src)
}

// ReplacePathInvalidCharacterToUnderline replace path invalid char to underline
// Deprecated
func ReplacePathInvalidCharacterToUnderline(s string) string {
	character := ReplacePathInvalidCharacter(s, "_")
	for {
		if !strings.Contains(character, "__") {
			break
		}
		character = strings.ReplaceAll(character, "__", "_")
	}
	return character
}

// RemovePathInvalidCharacter remove path invalid character
// Deprecated
func RemovePathInvalidCharacter(s string) string {
	return ReplacePathInvalidCharacter(s, "")
}

// ReplacePathInvalidCharacter replace path invalid char
// Deprecated
func ReplacePathInvalidCharacter(s string, replaceChar string) string {
	s = strings.ReplaceAll(s, `:`, replaceChar)
	s = strings.ReplaceAll(s, `*`, replaceChar)
	s = strings.ReplaceAll(s, `?`, replaceChar)
	s = strings.ReplaceAll(s, `<`, replaceChar)
	s = strings.ReplaceAll(s, `>`, replaceChar)
	s = strings.ReplaceAll(s, `"`, replaceChar)
	s = strings.ReplaceAll(s, `'`, replaceChar)
	s = strings.ReplaceAll(s, `|`, replaceChar)
	return s
}
