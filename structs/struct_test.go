package structs

import (
	"reflect"
	"testing"
)

func TestToFieldValueMap(t *testing.T) {
	type args struct {
		i interface{}
	}
	type test struct {
		Aaa string
		Bbb int
	}
	a := test{
		Aaa: "111",
		Bbb: 222,
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			args: args{i: a},
			want: map[string]interface{}{"Aaa": "111", "Bbb": 222},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFieldValueMap(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToFieldValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
