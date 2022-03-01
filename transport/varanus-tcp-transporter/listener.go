package varanustcptransporter

import (
	"net"

	varanus_core "github.com/Dione-Software/varanus-prototype/varanus-core"
	ma "github.com/multiformats/go-multiaddr"
)

type TcpListener struct {
	localMultiAddress *ma.Multiaddr
	internal net.TCPListener
}


func (l *TcpListener) LocalAddress() *ma.Multiaddr {
	return l.localMultiAddress
}

func (l *TcpListener) Accept() (varanus_core.TransportConnection, error) {
	tcpConn, err := l.internal.AcceptTCP()
	if err != nil {
		return nil, err
	}	
	remote := tcpConn.LocalAddr()
	remoteMultiaddr := AddrToMultiaddr(remote)
	conn := TcpConnection {
		internal: tcpConn,
		localAdrr: l.LocalAddress(),
		remoteAddr: &remoteMultiaddr,
	}
	return &conn, nil
}

func NewTcpListener(localAddress ma.Multiaddr) (*TcpListener, error) {
	tcpAddr, err := MultiaddrToTcpAddr(localAddress) 
	if err != nil {
		return nil, err
	}
	internalListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	ret := &TcpListener {
		localMultiAddress: &localAddress,
		internal: *internalListener,
	}
	return ret, nil
}
