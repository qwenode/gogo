package sanitize

import "testing"

func TestInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			args: args{
				str: "wfew123",
			},
			want: 123,
		},
		{
			args: args{
				str: "wfew1eee",
			},
			want: 1,
		},
		{
			args: args{
				str: "wf32w23gf",
			},
			want: 3223,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.str); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{str: "wie32w.ef0f3"},
			want: 32.03,
		},
		{
			args: args{str: "wie32w.ef0.f3"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.str); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}
