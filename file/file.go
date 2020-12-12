package file

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
)

// GetContents get file content
func GetContents(filename string) string {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(file)
}

// Exist check file or dir if exists
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
func Size(filename string) int64 {
	f, err := os.Stat(filename)
	if err != nil {
		return 0
	}
	return f.Size()
}

// WriteFileAppend Append content to the end of the file
func WriteFileAppend(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

// PutContents put content to the file,will clean old data
func PutContents(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(content)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return err
}

// Sha1 get file sha1 hash
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

//GetSize get file size
func GetSize(fileName string) int64 {
	stat, err := os.Stat(fileName)
	if err != nil {
		return 0
	}
	return stat.Size()
}

// IsDirectory check file is directory
func IsDirectory(filename string) bool {
	stat, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return stat.IsDir()
}
