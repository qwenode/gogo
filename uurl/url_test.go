package uurl

import "testing"

func TestUrlDecode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{s: "http%3A%2F%2Fwww.google.com"},
			want: "http://www.google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UrlDecode(tt.args.s); got != tt.want {
				t.Errorf("UrlDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
