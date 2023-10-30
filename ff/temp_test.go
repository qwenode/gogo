package ff

import (
	"log"
	"os"
	"testing"
)

func TestCreateTempFileWithString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *os.File
		wantErr bool
	}{
		{
			args: args{str: "哈哈哈"},
			//want: &os.File{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTempFileWithString(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTempFileWithString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println("CREATED TEMP FILE:", got.Name())
			got.Close()
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CreateTempFileWithString() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
