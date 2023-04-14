package ratelimit

import (
    "reflect"
    "testing"
    "time"
)

func TestNew(t *testing.T) {
    type args struct {
        limit int
    }
    tests := []struct {
        name string
        args args
        want *rateLimit
    }{
        {
            args: args{limit: 10},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := NewRateLimit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
                for i := 0; i < 20; i++ {
                    got.Add()
                    go func() {
                        defer got.Done()
                        time.Sleep(time.Second)
                        return
                    }()
                }
                got.Wait()
                //t.Errorf("New() = %v, want %v", got, tt.want)
            }
        })
    }
}
