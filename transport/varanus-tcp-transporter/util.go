package varanustcptransporter

import (
	"fmt"
	"net"
	"strconv"

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
		baseComponent, err := multiaddr.NewComponent("tcp", fmt.Sprint(port))
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


func MultiaddrToTcpAddr(address multiaddr.Multiaddr) (*net.TCPAddr, error) {
	stringAddress := ""
	stringAddressSet := false
	port := 0
	portSet := false
	shouldAbort := false
	var funcError error
	multiaddr.ForEach(address, func(c multiaddr.Component) bool {
		protocol := c.Protocol()
		switch protocol.Code {
		case multiaddr.P_TCP:
			portString := c.Value()
			port64, err := strconv.ParseInt(portString, 10, 64)
			if err != nil {
				shouldAbort = true
				funcError = err
				return false
			}
			port = int(port64)
			portSet = true
		case multiaddr.P_DNS:
			stringAddressSet = true
			stringAddress = c.Value()
		case multiaddr.P_IP4:
			stringAddressSet = true
			stringAddress = c.Value()
		case multiaddr.P_DNS4:
			stringAddressSet = true
			stringAddress = c.Value()
		case multiaddr.P_DNS6:
			stringAddressSet = true
			stringAddress = c.Value()
		case multiaddr.P_DNSADDR:
			stringAddressSet = true
			stringAddress = c.Value()
		case multiaddr.P_IP6:
			stringAddressSet = true
			stringAddress = c.Value()
		}
		return true
	})
	if shouldAbort {
		return nil, funcError
	}
	if !(stringAddressSet && portSet) {
		return nil, NotAllNecessarycomponents
	}
	stringAddress = fmt.Sprintf("%v:%v", stringAddress, port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", stringAddress)
	if err != nil {
		return nil, err
	}
	return tcpAddr, nil
}
