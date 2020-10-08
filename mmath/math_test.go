package mmath

import "testing"

func TestCeilInt(t *testing.T) {
	type args struct {
		first  int
		second int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			args: args{
				first:  0,
				second: 0,
			},
			want: 0,
		},
		{
			args: args{
				first:  0,
				second: 1,
			},
			want: 0,
		},
		{
			args: args{
				first:  1,
				second: 0,
			},
			want: 1,
		},
		{
			args: args{
				first:  10,
				second: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CeilInt(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CeilInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeilInt64(t *testing.T) {
	type args struct {
		first  int64
		second int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				first:  0,
				second: 0,
			},
			want: 0,
		},
		{
			args: args{
				first:  0,
				second: 1,
			},
			want: 0,
		},
		{
			args: args{
				first:  1,
				second: 0,
			},
			want: 1,
		},
		{
			args: args{
				first:  10,
				second: 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CeilInt64(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("CeilInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
