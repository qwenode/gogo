package file

import (
	"crypto/sha1"
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

func Sha1(fileName string) (string, error) {
	f, err := os.Open(fileName)
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
