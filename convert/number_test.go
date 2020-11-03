package convert

import "testing"

func TestToInt(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{str: "111"},
			want: 111,
		},
		{
			args: args{str: "x111"},
			want: 0,
		},
		{
			args: args{str: uint(1)},
			want: 1,
		},
		{
			args: args{str: int(1)},
			want: 1,
		},
		{
			args: args{str: int64(1)},
			want: 1,
		},
		{
			args: args{str: uint64(1)},
			want: 1,
		},
		{
			args: args{str: 0.33},
			want: 0,
		},
		{
			args: args{str: 1.33},
			want: 1,
		},
		{
			args: args{str: 1.55},
			want: 1,
		},
		{
			args: args{str: int64(22)},
			want: 22,
		},
		{
			args: args{str: float32(22.2)},
			want: 22,
		},
		{
			args: args{str: 1},
			want: 1,
		},
		{
			args: args{str: args{str: "x"}},
			want: 0,
		},
		{
			args: args{str: true},
			want: 1,
		},
		{
			args: args{str: false},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.str); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{str: "111"},
			want: 111,
		},
		{
			args: args{str: "x111"},
			want: 0,
		},
		{
			args: args{str: 0.33},
			want: 0,
		},
		{
			args: args{str: 1.33},
			want: 1,
		},
		{
			args: args{str: 1.55},
			want: 1,
		},
		{
			args: args{str: uint(1)},
			want: 1,
		},
		{
			args: args{str: int(1)},
			want: 1,
		},
		{
			args: args{str: int64(1)},
			want: 1,
		},
		{
			args: args{str: uint64(1)},
			want: 1,
		},
		{
			args: args{str: int64(22)},
			want: 22,
		},
		{
			args: args{str: float32(22.2)},
			want: 22,
		},
		{
			args: args{str: 1},
			want: 1,
		},
		{
			args: args{str: args{str: "x"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.str); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUInt(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{str: "111"},
			want: 111,
		},
		{
			args: args{str: "x111"},
			want: 0,
		},
		{
			args: args{str: 0.33},
			want: 0,
		},
		{
			args: args{str: 1.33},
			want: 1,
		},
		{
			args: args{str: 1.55},
			want: 1,
		},
		{
			args: args{str: int64(22)},
			want: 22,
		},
		{
			args: args{str: float32(22.2)},
			want: 22,
		},
		{
			args: args{str: uint(1)},
			want: 1,
		},
		{
			args: args{str: int(1)},
			want: 1,
		},
		{
			args: args{str: int64(1)},
			want: 1,
		},
		{
			args: args{str: uint64(1)},
			want: 1,
		},
		{
			args: args{str: args{str: "x"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUInt(tt.args.str); got != tt.want {
				t.Errorf("ToUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUInt64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{str: "111"},
			want: 111,
		},
		{
			args: args{str: "x111"},
			want: 0,
		},
		{
			args: args{str: 0.33},
			want: 0,
		},
		{
			args: args{str: 1.33},
			want: 1,
		},
		{
			args: args{str: 1.55},
			want: 1,
		},
		{
			args: args{str: int64(22)},
			want: 22,
		},
		{
			args: args{str: float32(22.2)},
			want: 22,
		},
		{
			args: args{str: uint(1)},
			want: 1,
		},
		{
			args: args{str: int(1)},
			want: 1,
		},
		{
			args: args{str: int64(1)},
			want: 1,
		},
		{
			args: args{str: uint64(1)},
			want: 1,
		},
		{
			args: args{str: args{str: "x"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUInt64(tt.args.str); got != tt.want {
				t.Errorf("ToUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		str interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{str: "111"},
			want: 111,
		},
		{
			args: args{str: "x111"},
			want: 0,
		},
		{
			args: args{str: 0.33},
			want: 0.33,
		},
		{
			args: args{str: 1.33},
			want: 1.33,
		},
		{
			args: args{str: 1.55},
			want: 1.55,
		},
		{
			args: args{str: int64(22)},
			want: 22,
		},
		{
			args: args{str: float64(22.21)},
			want: 22.21,
		},
		{
			args: args{str: float32(22.21)},
			want: 22.209999084472656,
		},
		{
			args: args{str: uint(1)},
			want: 1,
		},
		{
			args: args{str: int(1)},
			want: 1,
		},
		{
			args: args{str: int64(1)},
			want: 1,
		},
		{
			args: args{str: uint64(1)},
			want: 1,
		},
		{
			args: args{str: args{str: "x"}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.str); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
