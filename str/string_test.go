package str

import "testing"

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
