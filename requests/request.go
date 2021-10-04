package requests

import (
	"crypto/tls"
	"github.com/qwenode/gogo/convert"
	"net"
	"net/http"
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
