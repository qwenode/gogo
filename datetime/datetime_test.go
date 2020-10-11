package datetime

import "testing"

func TestGetUnixTime(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			want: GetUnixTime(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUnixTime(); got != tt.want {
				t.Errorf("GetUnixTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			want: GetCurrentTime(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentTime(); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBeginOfTheDayByInt64(t *testing.T) {
	type args struct {
		unix int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{unix: 1602426827},
			want: GetBeginOfTheDayByInt64(1602426827),
		}, {
			args: args{unix: 1602313200},
			want: GetBeginOfTheDayByInt64(1602313200),
		},
		{
			args: args{unix: 12259200},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBeginOfTheDayByInt64(tt.args.unix); got != tt.want {
				t.Errorf("GetBeginOfTheDayByInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBeginOfTheDayByInt(t *testing.T) {
	type args struct {
		unix int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{unix: 1602426827},
			want: GetBeginOfTheDayByInt(1602426827),
		}, {
			args: args{unix: 1602313200},
			want: GetBeginOfTheDayByInt(1602313200),
		},
		{
			args: args{unix: 12259200},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBeginOfTheDayByInt(tt.args.unix); got != tt.want {
				t.Errorf("GetBeginOfTheDayByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEndOfTheDayByInt64(t *testing.T) {
	type args struct {
		unix int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{unix: 1602426827},
			want: GetEndOfTheDayByInt64(1602426827),
		}, {
			args: args{unix: 1602313200},
			want: GetEndOfTheDayByInt64(1602313200),
		},
		{
			args: args{unix: 12259200},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEndOfTheDayByInt64(tt.args.unix); got != tt.want {
				t.Errorf("GetEndOfTheDayByInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEndOfTheDayByInt(t *testing.T) {
	type args struct {
		unix int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{unix: 1602426827},
			want: GetEndOfTheDayByInt(1602426827),
		}, {
			args: args{unix: 1602313200},
			want: GetEndOfTheDayByInt(1602313200),
		},
		{
			args: args{unix: 12259200},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEndOfTheDayByInt(tt.args.unix); got != tt.want {
				t.Errorf("GetEndOfTheDayByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
