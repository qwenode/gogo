package system

import "testing"

func TestGoroutineID(t *testing.T) {
	id := GoroutineID()
	tests := []struct {
		name string
		want int
	}{
		{
			want: id,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := id; got != tt.want {
				t.Errorf("GoroutineID() = %v, want %v", got, tt.want)
			}
		})
	}
}
