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
