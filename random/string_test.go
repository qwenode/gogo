package random

import "testing"

func TestGetRandString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 21},
			want: 21,
		},
		{
			args: args{n: 40},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandString(tt.args.n); len(got) != tt.want {
				t.Errorf("GetRandString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandStringNormal(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 21},
			want: 21,
		},
		{
			args: args{n: 40},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandStringNormal(tt.args.n); len(got) != tt.want {
				t.Errorf("GetRandStringNormal() = %v, want %v", got, tt.want)
			}
		})
	}
}
