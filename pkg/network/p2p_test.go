package network

import "testing"

func TestNewHost(t *testing.T) {
	cfg := &P2PConfig{ListenPort: 1, BootstrapPeers: []string{"a"}}
	h, err := NewHost(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if len(h.Peers) != 1 || h.Peers[0] != "a" {
		t.Fatalf("peers not set")
	}
}
