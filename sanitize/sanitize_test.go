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

func TestAlphabet(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "HKD 39.79+"},
			want: "HKD",
		},
		{
			args: args{str: "HKD 39.79+ test"},
			want: "HKD test",
		},
		{
			args: args{str: "   "},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alphabet(tt.args.str); got != tt.want {
				t.Errorf("Alphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphabetWithoutSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "HKD 39.79+"},
			want: "HKD",
		},
		{
			args: args{str: "HKD 39.79+ test"},
			want: "HKDtest",
		},
		{
			args: args{str: "HKD t39.tt79+ test"},
			want: "HKDttttest",
		},
		{
			args: args{str: "  "},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphabetWithoutSpace(tt.args.str); got != tt.want {
				t.Errorf("AlphabetWithoutSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHostName(t *testing.T) {
	type args struct {
		u string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{u: "http://www.amazon.com/"},
			want: "www.amazon.com",
		},
		{
			args: args{u: "http://www.amazon.com:883/xxxx"},
			want: "www.amazon.com",
		},
		{
			args: args{u: "https://amazon.com:333/"},
			want: "amazon.com",
		},
		{
			args: args{u: "www.amazon.com:883/xxx"},
			want: "www.amazon.com",
		},
		{
			args: args{u: "aaaa@www.amazon.com:883/xxx"},
			want: "www.amazon.com",
		},
		{
			args: args{u: " aaaa@www.amazon.com:883/xxx "},
			want: "www.amazon.com",
		},
		{
			args: args{u: "ftp://aaaa@www.amazon.com:883/xxx?xx=1"},
			want: "www.amazon.com",
		},
		{
			args: args{u: "://xx"},
			want: "",
		},
		{
			args: args{u: "xxxxx"},
			want: "xxxxx",
		},
		{
			args: args{u: "xxxxx<"},
			want: "xxxxx",
		},
		{
			args: args{u: "xxxxx)$"},
			want: "xxxxx",
		},
		{
			args: args{u: "1.1.1.1:22"},
			want: "1.1.1.1",
		},
		{
			args: args{u: "*.com.xx"},
			want: "*.com.xx",
		},
		{
			args: args{u: "*.*.com.xx"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HostName(tt.args.u); got != tt.want {
				t.Errorf("HostName() = %v, want %v", got, tt.want)
			}
		})
	}
}
