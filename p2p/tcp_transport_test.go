package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NOPHandshakeFunc,
		Decoder:       GOBDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)

	// assert.Equal(t, tr., listenAddress)

	// Server
	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
