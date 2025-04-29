package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	listenAddress := ":4000"
	transport := NewTCPTransport(listenAddress)
	if transport.listenAddress != listenAddress {
		t.Errorf("Expected listenAddress to be %s, got %s", listenAddress, transport.listenAddress)
	}
	if transport.peers == nil {
		t.Errorf("Expected peers to be initialized, got nil")
	}

	assert.NoError(t, transport.ListenAndAccept())

}
