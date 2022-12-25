package cmdline

import (
	"log"
	"testing"
)

func TestCommandLineFn(t *testing.T) {
	type args struct {
		fn commandFunc
		c  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{fn: func(output string, errCode int) bool {
				log.Println(output, errCode)
				return true
			}, c: "ping -c 1 127.0.0.1"},
			want: true,
		},
		{
			args: args{fn: func(output string, errCode int) bool {
				log.Println(output, errCode)
				return true
			}, c: "ping"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandFn(tt.args.fn, tt.args.c); got != tt.want {
				t.Errorf("CommandLineFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
