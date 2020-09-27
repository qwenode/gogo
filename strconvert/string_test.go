package strconvert

import "testing"

func TestToString(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{i: 111},
			want: "111",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.i); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
