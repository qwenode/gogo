package mmap

import (
    "reflect"
    "testing"
)

func TestNewDuplicateChecker(t *testing.T) {
    type args struct {
        needLock bool
    }
    type testCase[T interface{ string | int | float64 | float32 }] struct {
        name string
        args args
        want *duplicateCheck[T]
    }
    tests := []testCase[string]{
        {args: args{needLock: false}},
        {args: args{needLock: true}},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewDuplicateChecker[string](tt.args.needLock); !reflect.DeepEqual(got, tt.want) {
                if got.Exist("xx") {
                    t.Errorf("NewDuplicateChecker().Exist() = %v, want %v", true, false)
                }
                got.Add("xx")
                if !got.Exist("xx") {
                    t.Errorf("NewDuplicateChecker().Exist() = %v, want %v", false, true)
                }
            }
        })
    }
}
