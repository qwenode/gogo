package cmd

import (
	"testing"
)

func TestCommandFn(t *testing.T) {
	type args struct {
		name string
		fn   CommandFunc
		arg  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				name: "ping",
				fn: func(output string, errCode int) bool {
					//log.Println(output, errCode)
					return true
				},
				arg: []string{"127.0.0.1", "-c", "1"},
			},
			want: true,
		},
		{
			args: args{
				name: "test",
				fn: func(output string, errCode int) bool {
					return false
				},
				arg: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandFn(tt.args.name, tt.args.fn, tt.args.arg...); got != tt.want {
				t.Errorf("CommandFn() = %v, want %v", got, tt.want)
			}
		})
	}
}
