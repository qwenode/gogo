package ffunc

import (
	"errors"
	"log"
	"testing"
)

func TestRetry(t *testing.T) {
	type args struct {
		callable func() error
		retries  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				callable: func() error {
					log.Println("test retry 1 time")
					return nil
				},
				retries: 10,
			},
		},
		{
			args: args{
				callable: func() error {
					log.Println("test retry 10 times")
					return errors.New("xxx")
				},
				retries: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Retry(tt.args.callable, tt.args.retries); (err != nil) != tt.wantErr {
				t.Errorf("Retry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
