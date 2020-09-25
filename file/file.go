package file

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func WriteFileAppend(filename string, c []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(c)
	if err1 := f.Close(); err == nil {
		err = err1
	}

	return nil
}

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
