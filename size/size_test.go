package size

import (
	"testing"
)

func TestToHumanReadableSize(t *testing.T) {
	type args struct {
		size      float64
		precision int
		unit      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				size:      1024,
				precision: 2,
				unit:      SIZE_UNIT_BYTE,
			},
			want: "1024.00B",
		},
		{
			args: args{
				size:      1024,
				precision: 2,
				unit:      SIZE_UNIT_KILO_BYTE,
			},
			want: "1.00KB",
		},
		{
			args: args{
				size:      1024,
				precision: 2,
				unit:      SIZE_UNIT_KILO_BYTE,
			},
			want: "1.00KB",
		},
		{
			args: args{
				size:      159632,
				precision: 2,
				unit:      SIZE_UNIT_KILO_BYTE,
			},
			want: "155.89KB",
		},
		{
			args: args{
				size:      159632,
				precision: 2,
				unit:      SIZE_UNIT_KILO_BYTE,
			},
			want: "155.89KB",
		},
		{
			args: args{
				size:      15963238,
				precision: 9,
				unit:      SIZE_UNIT_KILO_BYTE,
			},
			want: "15589.099609375KB",
		},
		{
			args: args{
				size:      15963238,
				precision: 2,
				unit:      SIZE_UNIT_MEGA_BYTE,
			},
			want: "15.22MB",
		},
		{
			args: args{
				size:      15963238,
				precision: 4,
				unit:      SIZE_UNIT_GIGA_BYTE,
			},
			want: "0.0149GB",
		},
		{
			args: args{
				size:      15963238,
				precision: 2,
				unit:      11,
			},
			want: "0.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToHumanReadableSize(tt.args.size, tt.args.precision, tt.args.unit); got != tt.want {
				t.Errorf("ToHumanReadableSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
