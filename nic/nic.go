package nic

import (
	"net"
)

// GetAllBindAddress get all bind interface address
func GetAllBindAddress() []net.IP {
	ipList, _ := net.InterfaceAddrs()
	list := make([]net.IP, 0)
	for _, addr := range ipList {
		ip := addr.(*net.IPNet).IP
		list = append(list, ip)
	}
	return list
}

// GetAllIPV4BindAddress get all ipv4 bind address
func GetAllIPV4BindAddress() []net.IP {
	all := GetAllBindAddress()
	list := make([]net.IP, 0)
	for _, ip := range all {
		if false == IsIPv4IntranetAddress(ip) {
			list = append(list, ip)
		}
	}
	return list
}

// IsIPv4IntranetAddress check is intranet address
func IsIPv4IntranetAddress(ip interface{}) bool {
	ipc, ok := ip.(net.IP)
	if !ok {
		a := ip.(string)
		ipc = net.ParseIP(a)
	}
	v4 := ipc.To4()
	if v4 == nil {
		return true
	}
	if v4.IsGlobalUnicast() == false {
		return true
	}
	if v4[0] == 10 {
		return true
	}
	if v4[0] == 192 && v4[1] == 168 {
		return true
	}
	if v4[0] == 172 && v4[1] >= 16 && v4[1] <= 31 {
		return true
	}
	return false
}

// IsValidIPAddress check given ip address is valid
func IsValidIPAddress(ip string) bool {
	ipc := net.ParseIP(ip)
	if ipc.To4() != nil {
		return true
	}
	if ipc.To16() != nil {
		return true
	}
	return false
}

// IsValidIPV4Address check given ipv4 address is valid
func IsValidIPV4Address(ip string) bool {
	ipc := net.ParseIP(ip)
	if ipc.To4() != nil {
		return true
	}
	return false
}

// IsValidIPV6Address check given ipv6 address is valid
func IsValidIPV6Address(ip string) bool {
	ipc := net.ParseIP(ip)
	if ipc.To16() != nil {
		return true
	}
	return false
}
