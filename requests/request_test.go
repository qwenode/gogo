package requests

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	type args struct {
		url    string
		toFile string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			args: args{
				url:    "https://www.ipip.net",
				toFile: "./download_testcase",
			},
			want: 1,
		},
		{
			args: args{
				url:    "https://www.baidu.com",
				toFile: "./download_testcase",
			},
			want: 1,
		},
		{
			args: args{
				url:    "https://www.ipip.net",
				toFile: "/dev/test",
			},
			want:    0,
			wantErr: true,
		},
		{
			args: args{
				url:    "http://www.ipip.netx:99",
				toFile: "./download_testcase",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Download(tt.args.url, tt.args.toFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got < tt.want {
				t.Errorf("Download() got = %v, want > %v", got, tt.want)
			}
		})
	}
	os.Remove("./download_testcase")
}

func TestGetHeaderSize(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{url: "http://myip.ipip.net/"},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHeaderSize(tt.args.url); got < tt.want {
				t.Errorf("GetHeaderSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
