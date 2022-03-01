package varanustcptransporter

import (
	"net"

	"github.com/multiformats/go-multiaddr"
)

// AddrToMultiaddr performs a lookup and is therefore not very efficient
func AddrToMultiaddr(addr net.Addr) multiaddr.Multiaddr {
	switch addr.Network() {
	case "tcp":
		tcpAddr, err := net.ResolveTCPAddr("tcp", addr.String())
		if err != nil {
			panic(err)
		}
		ip := tcpAddr.IP
		port := tcpAddr.Port
		baseComponent, err := multiaddr.NewComponent("tcp", string(port))
		if err != nil {
			panic(err)
		}
		ipLen := len(ip)
		switch ipLen {
		case 4:
			ip4Component, err := multiaddr.NewComponent("ip4", ip.String())
			if err != nil {
				panic(err)
			}
			return baseComponent.Encapsulate(ip4Component)
		case 16:
			ip6Component, err := multiaddr.NewComponent("ip6", ip.String())
			if err != nil {
				panic(err)
			}
			return baseComponent.Encapsulate(ip6Component)
		default:
			panic("Ip len makes no sense")
		}
	default:
		panic("")
	}
}
