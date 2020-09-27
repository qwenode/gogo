package strconvert

import "testing"

func TestToInt(t *testing.T) {
	type args struct {
		str string
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.str); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
