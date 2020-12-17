package requests

import (
	"github.com/qwenode/gogo/convert"
	"io"
	"net/http"
	"os"
)

// GetHeaderSize get remote file size without download file
func GetHeaderSize(url string) int64 {
	head, err := http.Head(url)
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
