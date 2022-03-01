package varanustcptransporter

import (
	"errors"
	"net"
	"sync/atomic"

	ma "github.com/multiformats/go-multiaddr"
)


type TcpConnection struct {
	internal *net.TCPConn
	localAdrr *ma.Multiaddr
	remoteAddr *ma.Multiaddr
}


func (c *TcpConnection) Read(p []byte) (n int, err error) {
	return c.internal.Read(p)
}

func (c *TcpConnection) Write(p []byte) (n int, err error) {
	return c.internal.Write(p)
}

func (c *TcpConnection) Close() error {
	return c.internal.Close()
}

func (c *TcpConnection) LocalAddress() *ma.Multiaddr {
	return c.localAdrr
}

func (c *TcpConnection) RemoteAddress() *ma.Multiaddr {
	return c.remoteAddr
}

type TcpWriter struct {
	internal *net.TCPConn
	localAddr *ma.Multiaddr
	peerAddr *ma.Multiaddr
	closed *uint32
}

var ClosedError = errors.New("Already closed")

func (w *TcpWriter) Write(p []byte) (n int, err error)  {
	closed := atomic.LoadUint32(w.closed)
	if closed > 0 {
		return 0, ClosedError
	}
	return w.internal.Write(p)
}

func (w *TcpWriter) Close() error {
	atomic.AddUint32(w.closed, 1)
	return w.internal.Close()
}

func (c *TcpWriter) LocalAddress() *ma.Multiaddr {
	return c.localAddr
}

func (c *TcpWriter) RemoteAddress() *ma.Multiaddr {
	return c.peerAddr
}

type TcpReader struct {
	internal *net.TCPConn
	localAddr *ma.Multiaddr
	peerAddr *ma.Multiaddr
	closed *uint32
}

func (r *TcpReader) LocalAddress() *ma.Multiaddr {
	return r.localAddr
}

func (r *TcpReader) RemoteAddress() *ma.Multiaddr {
	return r.peerAddr
}

func (r *TcpReader) Read(p []byte) (n int, err error) {
	closed := atomic.LoadUint32(r.closed)
	if closed > 0 {
		return 0, ClosedError
	}
	return r.internal.Read(p)
}

func (r *TcpReader) Close() error {
	atomic.AddUint32(r.closed, 1)
	return r.internal.Close()
}

