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

func TestHostSendReceive(t *testing.T) {
	h, _ := NewHost(&P2PConfig{})
	if err := h.Send("p1", []byte("hello")); err != nil {
		t.Fatal(err)
	}
	msgs := h.Received("p1")
	if len(msgs) != 1 || string(msgs[0]) != "hello" {
		t.Fatalf("unexpected messages: %v", msgs)
	}
	if m := h.Received("unknown"); len(m) != 0 {
		t.Fatalf("expected empty slice, got %v", m)
	}
}
