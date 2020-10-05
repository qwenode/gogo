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
				start: -1,
				end:   2,
			},
			want: "ab",
		},
		{
			args: args{
				str:   "abcdefg",
				start: -1,
				end:   12,
			},
			want: "abcdefg",
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
				end:   2,
			},
			want: "国字",
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
