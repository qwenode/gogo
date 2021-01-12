package cmd

import (
	"log"
	"strings"
	"testing"
)

func TestCommandFn(t *testing.T) {
	type args struct {
		name string
		fn   commandFunc
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
			if got := CommandFn(tt.args.fn, tt.args.name, tt.args.arg...); got != tt.want {
				t.Errorf("CommandFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandRealtimeStdout(t *testing.T) {
	type args struct {
		fn   commandStdoutFunc
		name string
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
				fn: func(output string, errCode int, done bool) bool {
					//log.Println(output, errCode)
					if done {
						log.Println("finish", output)
					} else {
						log.Println("running:", output)
					}
					return true
				},
				arg: []string{"127.0.0.1", "-c", "1"},
			},
			want: true,
		},
		{
			args: args{
				name: "ping",
				fn: func(output string, errCode int, done bool) bool {
					//log.Println(output, errCode)
					if strings.Contains(output, "statistics") {
						return false
					}
					if done {
						log.Println("finish", output)
					} else {
						log.Println("running:", output)
					}
					return true
				},
				arg: []string{"127.0.0.1", "-c", "1"},
			},
			want: false,
		},
		{
			args: args{
				name: "/",
				fn: func(output string, errCode int, finish bool) bool {
					return false
				},
				arg: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandRealtimeStdout(tt.args.fn, tt.args.name, tt.args.arg...); got != tt.want {
				t.Errorf("CommandRealtimeStdout() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			if got := CommandLineFn(tt.args.fn, tt.args.c); got != tt.want {
				t.Errorf("CommandLineFn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommandLineRealtimeStdout(t *testing.T) {
	type args struct {
		fn commandStdoutFunc
		c  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				fn: func(output string, errCode int, done bool) bool {
					//log.Println(output, errCode)
					if done {
						log.Println("finish", output)
					} else {
						log.Println("running:", output)
					}
					return true
				},
				c: "ping 127.0.0.1 -c 1",
			},
			want: true,
		},
		{
			args: args{
				fn: func(output string, errCode int, done bool) bool {
					//log.Println(output, errCode)
					if strings.Contains(output, "statistics") {
						return false
					}
					if done {
						log.Println("finish", output)
					} else {
						log.Println("running:", output)
					}
					return true
				},
				c: "ping 127.0.0.1 -c 1",
			},
			want: false,
		},
		{
			args: args{
				fn: func(output string, errCode int, finish bool) bool {
					return false
				},
				c: "/",
			},
			want: false,
		},
		{
			args: args{
				fn: func(output string, errCode int, finish bool) bool {
					return false
				},
				c: "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandLineRealtimeStdout(tt.args.fn, tt.args.c); got != tt.want {
				t.Errorf("CommandLineRealtimeStdout() = %v, want %v", got, tt.want)
			}
		})
	}
}
