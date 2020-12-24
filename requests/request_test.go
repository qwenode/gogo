package requests

import (
	"github.com/qwenode/gogo/convert"
	"log"
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
		{
			args: args{url: "http://myip.ipip.netx/"},
			want: 0,
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

func TestDownloadFn(t *testing.T) {
	type args struct {
		fn     downloadFunc
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
				fn: func(nowSize int64) {
					log.Println("download size:" + convert.ToString(nowSize))
				},
				url:    "http://ecms.phome.net/downcenter/empirecms/ecms75/tool/ecms_memberconnect_qq.zip",
				toFile: "./download_testcase",
			},
			want:    1,
			wantErr: false,
		},
		{
			args: args{
				fn: func(nowSize int64) {
					log.Println("download size:" + convert.ToString(nowSize))
				},
				url:    "http://ecms.phome/wewf",
				toFile: "./download_testcase",
			},
			want:    0,
			wantErr: true,
		},
		{
			args: args{
				fn: func(nowSize int64) {
					log.Println("download size:" + convert.ToString(nowSize))
				},
				url:    "http://ecms.phome/wewf",
				toFile: "./",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DownloadFn(tt.args.fn, tt.args.url, tt.args.toFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadFn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got < tt.want {
				t.Errorf("DownloadFn() got = %v, want %v", got, tt.want)
			}
		})
	}
	os.Remove("./download_testcase")
}
