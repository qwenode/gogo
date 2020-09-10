package nic

import (
	"net"
)

func GetAllBindAddress() []net.IP {
	ipList, _ := net.InterfaceAddrs()
	list := make([]net.IP, 0)
	//ipList = append(ipList)
	for _, addr := range ipList {
		ip := addr.(*net.IPNet).IP
		list = append(list, ip)
	}
	return list
	//fmt.Println(ipList)
}

func GetAllIPV4BindAddress() []net.IP {
	all := GetAllBindAddress()
	list := make([]net.IP, 0)
	for _, ip := range all {
		if IsIPv4IntranetAddress(ip) {
			continue
		}
		list = append(list, ip)
	}
	return list
}

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
