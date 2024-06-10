package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddress := ":4000"
	tr := NewTCPTransport(listenAddress)

	assert.Equal(t, tr.listenAddress, listenAddress)

	// Server
	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
