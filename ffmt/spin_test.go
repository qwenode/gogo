package ffmt

import "testing"

func TestSpinPrint(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{a: "xxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SpinPrint(tt.args.a)
		})
	}
}
