package hashes

import "testing"

func Test_sha1(t *testing.T) {
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
			want: "d033e22ae348aeb5660fc2140aec35850c4da997",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha1(tt.args.str); got != tt.want {
				t.Errorf("Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}
