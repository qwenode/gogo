package nic

import (
	"net"
	"reflect"
	"testing"
)

func TestGetAllBindAddress(t *testing.T) {
	tests := []struct {
		name string
		want []net.IP
	}{
		{
			want: GetAllBindAddress(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllBindAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllBindAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllIPV4BindAddress(t *testing.T) {
	tests := []struct {
		name string
		want []net.IP
	}{
		{
			want: GetAllIPV4BindAddress(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllIPV4BindAddress(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllIPV4BindAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv4IntranetAddress(t *testing.T) {
	type args struct {
		ip interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{ip: "127.0.0.1"},
			want: true,
		},
		{
			//name: "2",
			args: args{ip: "192.168.0.1"},
			want: true,
		},
		{
			//name: "3",
			args: args{ip: "223.2.1.2"},
			want: false,
		},
		{
			args: args{ip: "10.0.0.1"},
			want: true,
		},
		{
			args: args{ip: "172.16.0.1"},
			want: true,
		},
		{
			args: args{ip: "372.316.0.1"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv4IntranetAddress(tt.args.ip); got != tt.want {
				t.Errorf("IsIPv4IntranetAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
