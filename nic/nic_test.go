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

func TestIsValidIPAddress(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				ip: "1.1.1.1",
			},
			want: true,
		},
		{
			args: args{
				ip: "1.1.1.1111",
			},
			want: false,
		},
		{
			args: args{
				ip: "2001:0DB8:0000:0023:0008:0800:200C:417A",
			},
			want: true,
		},
		{
			args: args{
				ip: "2001:DB8:0:23:8:800:200C:417A",
			},
			want: true,
		},
		{
			args: args{
				ip: "2001:DB8:0:23:8:800:200C:417A1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIPAddress(tt.args.ip); got != tt.want {
				t.Errorf("IsValidIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIPV4Address(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				ip: "1.1.1.1",
			},
			want: true,
		},
		{
			args: args{
				ip: "1.1.1.1111",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIPV4Address(tt.args.ip); got != tt.want {
				t.Errorf("IsValidIPV4Address() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidIPV6Address(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				ip: "2001:DB8:0:23:8:800:200C:417A",
			},
			want: true,
		},
		{
			args: args{
				ip: "2001:DB8:0:23:8:800:200C:417A1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidIPV6Address(tt.args.ip); got != tt.want {
				t.Errorf("IsValidIPV6Address() = %v, want %v", got, tt.want)
			}
		})
	}
}
