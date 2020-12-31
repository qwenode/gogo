package requests

import (
	"crypto/tls"
	"github.com/qwenode/gogo/convert"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

// GetHeaderSize get remote file size without download file
func GetHeaderSize(url string) int64 {
	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 5 * time.Second,
			}).DialContext,
			IdleConnTimeout:       5 * time.Second,
			TLSHandshakeTimeout:   5 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, _ := http.NewRequest(http.MethodHead, url, nil)
	head, err := client.Do(req)
	if err != nil {
		return 0
	}
	defer head.Body.Close()
	return convert.ToInt64(head.Header.Get("Content-Length"))
}

// Download save http response body to file,download large file
func Download(url string, toFile string) (int64, error) {
	create, err := os.Create(toFile)
	if err != nil {
		return 0, err
	}
	defer create.Close()
	get, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer get.Body.Close()
	written, _ := io.Copy(create, get.Body)
	return written, nil
}

// downloadFunc DownloadFn callback
type downloadFunc func(nowSize int64)

// DownloadFn save http response body to file,download large file
func DownloadFn(fn downloadFunc, url string, toFile string) (int64, error) {
	dst, err := os.Create(toFile)
	if err != nil {
		return 0, err
	}
	defer dst.Close()
	get, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer get.Body.Close()
	buf := make([]byte, 32*1024)
	written := int64(0)
	for {
		nr, er := get.Body.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
				fn(written)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	//written, _ := io.Copy(create, get.Body)
	//get.Body.Read()
	return written, nil
}
