package sanitize

import (
	"net/url"
	"testing"
)

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
		{
			args: args{str: "-fe1"},
			want: -1,
		},
		{
			args: args{str: "-fe0"},
			want: 0,
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
		{
			args: args{str: "wef|2"},
			want: "wef",
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

func TestAlphabetNumber(t *testing.T) {
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
			want: "HKD 3979",
		},
		{
			args: args{str: "HKD 39.79+ test"},
			want: "HKD 3979 test",
		},
		{
			args: args{str: "   "},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphabetNumber(tt.args.str); got != tt.want {
				t.Errorf("AlphabetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlphabetNumberWithoutSpace(t *testing.T) {
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
			want: "HKD3979",
		},
		{
			args: args{str: "HKD 39.79+ test"},
			want: "HKD3979test",
		},
		{
			args: args{str: "HKD t39.tt79+ test"},
			want: "HKDt39tt79test",
		},
		{
			args: args{str: "  "},
			want: "",
		},
		{
			args: args{str: "wef|2"},
			want: "wef2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AlphabetNumberWithoutSpace(tt.args.str); got != tt.want {
				t.Errorf("AlphabetNumberWithoutSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectoryPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{path: "/aaa/bb"},
			want: "/aaa/bb",
		},
		{
			args: args{path: "//wef"},
			want: "/wef",
		},
		{
			args: args{path: "///"},
			want: "/",
		},
		{
			args: args{path: "??"},
			want: "/",
		},
		{
			args: args{path: "//aaa///bb"},
			want: "/aaa/bb",
		},
		{
			args: args{path: "/.../aaa///bb"},
			want: "/aaa/bb",
		},
		{
			args: args{path: "/.../aa-_a///bb?"},
			want: "/aa-_a/bb",
		},
		{
			args: args{path: "/.../aaa/123//bb"},
			want: "/aaa/123/bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirectoryPath(tt.args.path); got != tt.want {
				t.Errorf("DirectoryPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleSpaceToSingle(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "a   b   c"},
			want: "a b c",
		},
		{
			args: args{str: "a   b   c   1 22 2 33"},
			want: "a b c 1 22 2 33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultipleSpaceToSingle(tt.args.str); got != tt.want {
				t.Errorf("MultipleSpaceToSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUInt(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
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
		{
			args: args{str: "-fe1"},
			want: 1,
		},
		{
			args: args{str: "-fe0"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UInt(tt.args.str); got != tt.want {
				t.Errorf("UInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{path: "/aaa/bb"},
			want: "/aaa/bb",
		},
		{
			args: args{path: "//wef"},
			want: "/wef",
		},
		{
			args: args{path: "///"},
			want: "/",
		},
		{
			args: args{path: "??"},
			want: "/",
		},
		{
			args: args{path: "//aaa///bb"},
			want: "/aaa/bb",
		},
		{
			args: args{path: "/.../aaa///bb"},
			want: "/.../aaa/bb",
		},
		{
			args: args{path: "/.../aaa///bb.php"},
			want: "/.../aaa/bb.php",
		},
		{
			args: args{path: "/.../aa-_a///bb?"},
			want: "/.../aa-_a/bb",
		},
		{
			args: args{path: "/.../aaa/123//bb"},
			want: "/.../aaa/123/bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilePath(tt.args.path); got != tt.want {
				t.Errorf("FilePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripHtml(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "<aa>"},
			want: "aa",
		},
		{
			args: args{str: url.PathEscape("<aa></bb>")},
			want: "aa%2Fbb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripHtml(tt.args.str); got != tt.want {
				t.Errorf("StripHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveAllSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "aa ,b b"},
			want: "aa,bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAllSpace(tt.args.str); got != tt.want {
				t.Errorf("RemoveAllSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultipleBackslashToSingle(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{str: "\\\\\\\\\\\\\\\\\\\\\\"},
			want: "\\",
		},
		{
			args: args{str: "\\\\\\\\\\\\\\\\\\\\"},
			want: "\\",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultipleBackslashToSingle(tt.args.str); got != tt.want {
				t.Errorf("MultipleBackslashToSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}
