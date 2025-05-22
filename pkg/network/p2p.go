// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package network

import "sync"

// P2PConfig holds peer configuration

type P2PConfig struct {
	ListenPort     int
	BootstrapPeers []string
}

// Host represents a minimal P2P node.
type Host struct {
	cfg      *P2PConfig
	Peers    []string
	mu       sync.Mutex
	messages map[string][][]byte
}

// NewHost creates a new host using the given configuration.
func NewHost(cfg *P2PConfig) (*Host, error) {
	return &Host{cfg: cfg, Peers: cfg.BootstrapPeers, messages: make(map[string][][]byte)}, nil
}

// Send stores the message for the given peer.
func (h *Host) Send(peer string, msg []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.messages[peer] = append(h.messages[peer], append([]byte(nil), msg...))
	return nil
}

// Received returns all messages sent to the given peer.
func (h *Host) Received(peer string) [][]byte {
	h.mu.Lock()
	defer h.mu.Unlock()
	msgs := h.messages[peer]
	out := make([][]byte, len(msgs))
	for i, m := range msgs {
		out[i] = append([]byte(nil), m...)
	}
	return out
}
