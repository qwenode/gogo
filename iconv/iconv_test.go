package iconv

import "testing"

func TestUTF8ToGBKString(t *testing.T) {
	type args struct {
		str string
	}
	gbkString, _ := UTF8ToGBKString("这是UTF8中文")
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{
			args:  args{str: "这是UTF8中文"},
			wantS: gbkString,
		},
		{
			args:  args{str: ""},
			wantS: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := UTF8ToGBKString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("UTF8ToGBKString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("UTF8ToGBKString() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestGBKToUTF8String(t *testing.T) {
	type args struct {
		str string
	}
	gbkStr, _ := UTF8ToGBKString("这是GBK转UTF8")
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{
			args:  args{str: gbkStr},
			wantS: "这是GBK转UTF8",
		},
		{
			args:  args{str: ""},
			wantS: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := GBKToUTF8String(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("GBKToUTF8String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("GBKToUTF8String() gotS = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
