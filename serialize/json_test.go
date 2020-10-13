package serialize

import "testing"

func TestJsonEncode(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{v: "a"},
			want: "\"a\"",
		},
		{
			args: args{v: map[string]string{"a": "b"}},
			want: "{\"a\":\"b\"}",
		},
		{
			args: args{v: make(chan int)},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonEncode(tt.args.v); got != tt.want {
				t.Errorf("JsonEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
