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
		filename string
	}
	sha1, _ := Sha1("file.go")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{filename: "file.go"},
			want: sha1,
		},
		{
			args:    args{filename: "."},
			want:    "",
			wantErr: true,
		},
		{
			args:    args{filename: "not_exists_file.noexists"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sha1(tt.args.filename)
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

func TestSha256(t *testing.T) {
	type args struct {
		filename string
	}
	sha1, _ := Sha256("file.go")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{filename: "file.go"},
			want: sha1,
		},
		{
			args:    args{filename: "."},
			want:    "",
			wantErr: true,
		},
		{
			args:    args{filename: "not_exists_file.noexists"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sha256(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sha256() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sha256() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	type args struct {
		filename string
	}
	sha1, _ := Md5("file.go")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{filename: "file.go"},
			want: sha1,
		},
		{
			args:    args{filename: "."},
			want:    "",
			wantErr: true,
		},
		{
			args:    args{filename: "not_exists_file.noexists"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Md5(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Md5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Md5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrc32(t *testing.T) {
	type args struct {
		filename string
	}
	sha1, _ := Crc32("file.go")
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			args: args{filename: "file.go"},
			want: sha1,
		},
		{
			args:    args{filename: "."},
			want:    "",
			wantErr: true,
		},
		{
			args:    args{filename: "not_exists_file.noexists"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Crc32(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crc32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Crc32() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExist(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{filename: "file.go"},
			want: true,
		},
		{
			args: args{filename: "fff.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exist(tt.args.filename); got != tt.want {
				t.Errorf("Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
