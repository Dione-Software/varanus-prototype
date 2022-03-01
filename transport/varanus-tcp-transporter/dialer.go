package varanustcptransporter

import (
	ma "github.com/multiformats/go-multiaddr"
	varanus_core "github.com/Dione-Software/varanus-prototype/varanus-core"
)

type TcpDialer struct {}



func (d *TcpDialer) Dial(address *ma.Multiaddr) (varanus_core.TransportConnection, error) {
	panic("not implemented") // TODO: Implement
}

