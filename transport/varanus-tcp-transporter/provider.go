package varanustcptransporter

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	varanus_core "github.com/Dione-Software/varanus-prototype/varanus-core"
	ma "github.com/multiformats/go-multiaddr"
)


type TcpProvider struct {}

func leel(member varanus_core.TransportProvider) {
}


var tcpProtocol = ma.ProtocolWithCode(ma.P_TCP)

func (p *TcpProvider) ProtocolIdentifier() *ma.Protocol {
	return &tcpProtocol
}

// ContainsNecessaryProtocols returns true, if the given Multiaddress
// contains all necessary protocols to call on this provider with
// this Multiaddress
func (p *TcpProvider) ContainsNecessaryProtocols(address *ma.Multiaddr) bool {
	ret, _, _ := ContainsNecessaryProtocols(address)
	return ret
}

// Dialer returns the generic associated dialer
func (p *TcpProvider) Dialer(localAddress ma.Multiaddr) (varanus_core.TransportDialer, error) {
	return NewTcpDialer(localAddress)
}

// Listener returns the generic associated listener
func (p *TcpProvider) Listener(localAddress ma.Multiaddr) (varanus_core.TransportListener, error) {
	return NewTcpListener(localAddress)
}

var NotAllNecessarycomponents = errors.New("Given address didn't contain all necessary components")

func ContainsNecessaryProtocols(address *ma.Multiaddr) (bool, *net.TCPAddr, error) {
	addressProvided := false
	var addressString string
	tcpPortProvided := false
	var tcpPort int
	var err error
	testFunction := func (c ma.Component) bool {
		protocol := c.Protocol()
		switch protocol.Code {
		case ma.P_IP4:
			addressProvided = true 
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_DNS:
			addressProvided = true
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_DNS4:
			addressProvided = true
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_DNS6:
			addressProvided = true
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_DNSADDR:
			addressProvided = true
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_IP6:
			addressProvided = true
			addressString, err = c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
		case ma.P_TCP:
			tcpPortProvided = true
			portString, err := c.ValueForProtocol(protocol.Code)
			if err != nil {
				panic(err)
			}
			tmp, err := strconv.ParseInt(portString, 10, 64)
			tcpPort = int(tmp)
		}
		return true
	}
	ma.ForEach(*address, testFunction)
	valid := addressProvided && tcpPortProvided
	if !valid {
		return false, nil, NotAllNecessarycomponents 
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%v:%v", addressString, tcpPort))
	if err != nil {
		return false, nil, err
	}
	return addressProvided && tcpPortProvided, tcpAddr, nil
}
