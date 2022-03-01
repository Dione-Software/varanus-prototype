package varanustcptransporter

import (
	"net"

	varanus_core "github.com/Dione-Software/varanus-prototype/varanus-core"
	ma "github.com/multiformats/go-multiaddr"
)

type TcpDialer struct {
	localAddress *net.TCPAddr
	localMultiaddress ma.Multiaddr
}

func NewTcpDialer(localAddress ma.Multiaddr) (*TcpDialer, error) {
	localNetAddress, err := MultiaddrToTcpAddr(localAddress)
	if err != nil {
		return nil, err
	}
	ret := &TcpDialer{
		localAddress: localNetAddress,
		localMultiaddress: localAddress,
	}
	return ret, nil
}

func (d *TcpDialer) Dial(address ma.Multiaddr) (varanus_core.TransportConnection, error) {
	remoteAddress, err := MultiaddrToTcpAddr(address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", d.localAddress, remoteAddress)
	if err != nil {
		return nil, err
	}
	ret := &TcpConnection{
		internal: conn,
		localAdrr: &d.localMultiaddress,
		remoteAddr: &address,
	}
	return ret, nil
}

