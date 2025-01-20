package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	opts := TCPTransportOpts{
		ListenAddress:    listenAddr,
		HandshakeHandler: NOPHandshakeFunc,
	}
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddress, listenAddr)
}
