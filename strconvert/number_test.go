package strconvert

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.str); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
