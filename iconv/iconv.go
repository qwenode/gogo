package iconv

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// GBKToUTF8 convert gbk to utf8 bytes
func GBKToUTF8(inByte []byte) (b []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(inByte), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(reader)
	return
}

// UTF8ToGBK convert utf8 to gbk bytes
func UTF8ToGBK(inByte []byte) (b []byte, err error) {
	reader := transform.NewReader(bytes.NewReader(inByte), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(reader)
	return
}

// GBKToUTF8String convert gbk to utf8 string
func GBKToUTF8String(str string) (s string, err error) {
	r, err := GBKToUTF8([]byte(str))
	return string(r), err
}

// UTF8ToGBKString convert utf8 to gbk string
func UTF8ToGBKString(str string) (s string, err error) {
	r, err := UTF8ToGBK([]byte(str))
	return string(r), err
}
