package requests

import (
	"testing"
)


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
			args: args{url: "https://ifconfig.io/"},
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
