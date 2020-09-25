package file

import (
	"os"
	"testing"
)

func TestWriteFileAppend(t *testing.T) {
	type args struct {
		filename string
		c        []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				filename: "test.txt",
				c:        []byte("haha"),
			},
			wantErr: false,
		},
		{
			args: args{
				filename: ".",
				c:        []byte("haha"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFileAppend(tt.args.filename, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("WriteFileAppend() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.Remove("test.txt")
		})
	}
}
