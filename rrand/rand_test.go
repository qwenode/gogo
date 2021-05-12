package rrand

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestIntn(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 20},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := 0
			if got = Intn(tt.args.n); got > tt.want {
				t.Errorf("Intn() = %v, want %v", got, tt.want)
			}
			log.Println("rand got:", got)
		})
	}
}

func TestInt63(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int63(); got < 0 {
				t.Errorf("Int63() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(); got < 0 {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt31(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int31(); got < 0 {
				t.Errorf("Int31() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt63n(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{n: 111},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int63n(tt.args.n); got < 0 {
				t.Errorf("Int63n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt31n(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{n: 10},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int31n(tt.args.n); got < 0 {
				t.Errorf("Int31n() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		want uint32
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(); got < 0 {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(); got < 0 {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(); got < 0 {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name string
		want float32
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(); got < 0 {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpFloat64(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpFloat64(); got < 0 {
				t.Errorf("ExpFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormFloat64(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormFloat64(); got < 0 {
				t.Errorf("NormFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	words := strings.Fields("ink runs from the corners of my mouth")
	Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}

func TestRead(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		{
			args:    args{p: []byte{1, 2, 3}},
			wantN:   3,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Read() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestPerm(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{n: 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perm(tt.args.n); len(got) != tt.want {
				t.Errorf("Perm() = %v, want %v", got, tt.want)
			}
		})
	}
}
