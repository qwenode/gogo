package array

import (
	"reflect"
	"testing"
)

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

func TestMergeIntRange(t *testing.T) {
	type args struct {
		data []IntegerRange
	}
	tests := []struct {
		name string
		args args
		want []IntegerRange
	}{
		{
			args: args{data: []IntegerRange{
				{80, 0}, {999, 0}, {80, 0},
				{80, 0}, {80, 0}, {80, 0},
				{80, 0},
			}},
			want: []IntegerRange{
				{80, 0}, {999, 0},
			},
		},
		{
			args: args{data: []IntegerRange{
				{Low: 0, High: 0}, {1, 0}, {22, 33},
				{23, 34}, {15, 32}, {15, 22},
			}},
			want: []IntegerRange{
				{Low: 0, High: 0}, {1, 0}, {15, 34},
			},
		},
		{
			args: args{data: []IntegerRange{}},
			want: nil,
		},
		{
			args: args{data: []IntegerRange{{12, 0}, {13, 0}}},
			want: []IntegerRange{{12, 0}, {13, 0}},
		},
		{
			args: args{data: []IntegerRange{{12, 0}}},
			want: []IntegerRange{{12, 0}},
		},
		{
			args: args{data: []IntegerRange{{80, 0}, {80, 90}, {80, 88}}},
			want: []IntegerRange{{80, 90}},
		},
		{
			args: args{data: []IntegerRange{{80, 88}, {80, 90}, {80, 88}}},
			want: []IntegerRange{{80, 90}},
		},
		{
			args: args{data: []IntegerRange{{80, 0}, {90, 0}}},
			want: []IntegerRange{{80, 0}, {90, 0}},
		},
		{
			args: args{data: []IntegerRange{{80, 0}, {0, 90}, {0, 0}}},
			want: []IntegerRange{{0, 90}},
		},
		{
			args: args{data: []IntegerRange{
				{Low: 0, High: 0}, {1, 0}, {22, 33},
				{23, 34}, {15, 32}, {15, 22},
				{0, 12},
			}},
			want: []IntegerRange{
				{Low: 0, High: 12}, {15, 34},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntRange(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
