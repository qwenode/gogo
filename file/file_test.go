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

func TestSha1(t *testing.T) {
	type args struct {
		fileName string
	}
	sha1, _ := Sha1("file.go")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{fileName: "file.go"},
			want: sha1,
		},
		{
			args:    args{fileName: "."},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sha1(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha1() got = %v, want %v", got, tt.want)
			}
		})
	}
}
