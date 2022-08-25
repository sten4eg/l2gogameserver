package config

import (
	"log"
	"net"
	"net/netip"
	"strings"
)

func GetSubNets() []string {
	var res []string
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return res
	}
	for _, v := range interfaces {
		addrs, err := v.Addrs()
		if err != nil {
			continue
		}
		for _, vv := range addrs {
			addr, err := netip.ParsePrefix(vv.String())
			if err != nil {
				continue
			}
			if addr.Addr().Is4() {
				res = append(res, addr.String())
			}
		}
	}
	return res
}

func GetHosts() []string {
	subNets := GetSubNets()
	var res []string
	for _, v := range subNets {
		b, _, ok := strings.Cut(v, "/")
		if ok {
			res = append(res, b)
		}
	}
	return res
}
