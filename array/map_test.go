package array

import (
	"testing"
)

func TestGetMapIntKeys(t *testing.T) {
	type args struct {
		m map[int]interface{}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{m: map[int]interface{}{0: "haha", 2: "ss", 3: "hehe"}},
			want: []int{0, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMapIntKeys(tt.args.m)
			for _, i := range got {
				if !IntInSlice(i, tt.want) {
					t.Errorf("GetMapIntKeys() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
