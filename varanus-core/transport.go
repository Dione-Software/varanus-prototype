package varanus_core

import (
	"io"
	ma "github.com/multiformats/go-multiaddr"
)

// TransportProvider is a generic interface for any kind of transport.
type TransportProvider interface {
	ProtocolIdentifier() *ma.Protocol
	// ContainsNecessaryProtocols returns true, if the given Multiaddress
	// contains all necessary protocols to call on this provider with
	// this Multiaddress
	ContainsNecessaryProtocols(address *ma.Multiaddr) bool
	// Dialer returns the generic associated dialer
	Dialer(localAddress ma.Multiaddr) (TransportDialer, error)
	// Listener returns the generic associated listener
	Listener(localAddress ma.Multiaddr) (TransportListener, error)
}

type TransportDialer interface {
	Dial(address ma.Multiaddr) (TransportConnection, error)
}

type TransportListener interface {
	LocalAddress() *ma.Multiaddr
	Accept() (TransportConnection, error)
}

type TransportConnection interface {
	io.ReadWriteCloser
	TransportConnectionMetadata
}

type TransportWriter interface {
	io.WriteCloser
	TransportConnectionMetadata
}

type TransportReader interface {
	io.ReadCloser
	TransportConnectionMetadata
}

type TransportConnectionMetadata interface {
	LocalAddress() *ma.Multiaddr
	RemoteAddress() *ma.Multiaddr
}
