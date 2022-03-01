package varanustcptransporter_test

import (
	"errors"
	"testing"

	varanustcptransporter "github.com/Dione-Software/varanus-prototype/transport/varanus-tcp-transporter"
	"github.com/multiformats/go-multiaddr"
)

func TestContainsNecessaryProtocols(t *testing.T) {
	address1, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/1234/udp/1333")
	if err != nil {
		t.Errorf("Error while parsing multiaddr %v", err)
	}
	contains, parsedAddress, err := varanustcptransporter.ContainsNecessaryProtocols(&address1)
	if err != nil {
		t.Errorf("Error while checking multiaddr %v", err)
	}
	if !contains {
		t.Errorf("Given multiaddress should contain all necessary information")
	}
	if parsedAddress.String() != "127.0.0.1:1234" {
		t.Errorf("Parsed address, was not expected %#v", parsedAddress)
	}
	address2, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/udp/1333")
	if err != nil {
		t.Errorf("Error while parsing multiaddr %v", err)
	}
	contains, parsedAddress, err = varanustcptransporter.ContainsNecessaryProtocols(&address2)
	if !errors.Is(err, varanustcptransporter.NotAllNecessarycomponents) {
		t.Errorf("The error message was unexpected %v", err)
	}
	if parsedAddress != nil {
		t.Error("Shouldn't return a tcp address")
	}
	if contains {
		t.Error("Wrong decision made")
	}
	address3, err := multiaddr.NewMultiaddr("/dns/www.google.com/tcp/80")
	if err != nil {
		t.Errorf("Error while parsing multiaddr %v", err)
	}
	contains, parsedAddress, err = varanustcptransporter.ContainsNecessaryProtocols(&address3)
}
