package hashes

import "testing"

func TestCrc32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "admin"},
			want: "880e0d76",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crc32(tt.args.str); got != tt.want {
				t.Errorf("Crc32() = %v, want %v", got, tt.want)
			}
		})
	}
}
