package array

import "testing"

func TestSliceContainInt(t *testing.T) {
	type args struct {
		search int
		slice  []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				search: 1,
				slice:  []int{1, 3, 5},
			},
			want: true,
		},
		{
			args: args{
				search: 1,
				slice:  []int{0, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntInSlice(tt.args.search, tt.args.slice); got != tt.want {
				t.Errorf("IntInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	type args struct {
		search string
		slice  []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				search: "1",
				slice:  []string{"1", "3", "5"},
			},
			want: true,
		},
		{
			args: args{
				search: "1",
				slice:  []string{"0", "3", "5"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringInSlice(tt.args.search, tt.args.slice); got != tt.want {
				t.Errorf("StringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
