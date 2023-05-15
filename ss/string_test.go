package ss

import (
	"reflect"
	"testing"
)

func TestSubstr(t *testing.T) {
	type args struct {
		str   string
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str:   "abcdefg",
				start: -10,
				end:   2,
			},
			want: "ab",
		},
		{
			args: args{
				str:   "abcdefg",
				start: 0,
				end:   2,
			},
			want: "ab",
		},
		{
			args: args{
				str:   "abcdefg",
				start: 1,
				end:   2,
			},
			want: "bc",
		},
		{
			args: args{
				str:   "abcdefg",
				start: 1,
				end:   -2,
			},
			want: "",
		},
		{
			args: args{
				str:   "abcdefg",
				start: -1,
				end:   2,
			},
			want: "g",
		},
		{
			args: args{
				str:   "abcdefg",
				start: -1,
				end:   12,
			},
			want: "g",
		},
		{
			args: args{
				str:   "abcdefg",
				start: 0,
				end:   7,
			},
			want: "abcdefg",
		},
		{
			args: args{
				str:   "中国字",
				start: 0,
				end:   7,
			},
			want: "中国字",
		},
		{
			args: args{
				str:   "中国字",
				start: 0,
				end:   1,
			},
			want: "中",
		},
		{
			args: args{
				str:   "国字",
				start: 1,
				end:   1,
			},
			want: "字",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substr(tt.args.str, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Substr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCutByEndString(t *testing.T) {
	type args struct {
		str    string
		endStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str:    "Pure organic linen fabrics, clothes and roses that last by LinenRoses",
				endStr: "by",
			},
			want: "Pure organic linen fabrics, clothes and roses that last",
		},
		{
			args: args{
				str:    "Pure organic linen fabrics, clothes and roses that last",
				endStr: "by",
			},
			want: "Pure organic linen fabrics, clothes and roses that last",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutByEndString(tt.args.str, tt.args.endStr); got != tt.want {
				t.Errorf("CutByEndString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastElemBySep(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "ab,bc",
				sep: ",",
			},
			want: "bc",
		},
		{
			args: args{
				str: "abbc",
				sep: ",",
			},
			want: "abbc",
		},
		{
			args: args{
				str: "ab,bc,cc",
				sep: ",",
			},
			want: "cc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastElemBySep(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("GetLastElemBySep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFirstElemBySep(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "ab,bc",
				sep: ",",
			},
			want: "ab",
		},
		{
			args: args{
				str: "abbc",
				sep: ",",
			},
			want: "abbc",
		},
		{
			args: args{
				str: "ab,bc,cc",
				sep: ",",
			},
			want: "ab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFirstElemBySep(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("GetFirstElemBySep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSecondElemBySep(t *testing.T) {
	type args struct {
		str string
		sep string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				str: "ab,bc",
				sep: ",",
			},
			want: "bc",
		},
		{
			args: args{
				str: "abbc",
				sep: ",",
			},
			want: "abbc",
		},
		{
			args: args{
				str: "ab,bc,cc",
				sep: ",",
			},
			want: "bc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSecondElemBySep(tt.args.str, tt.args.sep); got != tt.want {
				t.Errorf("GetSecondElemBySep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractUrls(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{str: "【淘宝】https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T CZ0001 「lulu短裙女运动裙裤跑步健身网球裙高腰速干羽毛球裙休闲高尔夫裙」\n点击链接直接打开 或者 淘宝搜索直接打开"},
			want: []string{"https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T"},
		},
		{
			args: args{str: ""},
			want: []string{},
		},
		{
			args: args{str: "【淘宝】https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T CZ0001 「lulu短裙女运动裙裤跑步健身网球裙高腰 https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T 速干羽毛球裙休闲高尔夫裙」\n点击链接直接打开 或者 淘宝搜索直接打开"},
			want: []string{"https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T", "https://m.tb.cn/h.U5y5PTs?tk=C7lddYenQ5T"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractUrls(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		str   string
		finds []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{str: "testxx1100022", finds: []string{"aa", "000"}},
			want: true,
		},
		{
			args: args{str: "testxx1100022", finds: []string{"aa", "bb"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.str, tt.args.finds...); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAllOfAll(t *testing.T) {
	type args struct {
		str   string
		finds []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				str:   "abc",
				finds: []string{"ABDD"},
			},
			want: false,
		},
		{
			args: args{
				str:   "abc",
				finds: []string{"abcda"},
			},
			want: true,
		},
		{
			args: args{
				str:   "abcda",
				finds: []string{"bcd"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAllOfAll(tt.args.str, tt.args.finds...); got != tt.want {
				t.Errorf("ContainsAllOfAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
