// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package network

// P2PConfig holds peer configuration

type P2PConfig struct {
	ListenPort     int
	BootstrapPeers []string
}

// Host represents a minimal P2P node.
type Host struct {
	cfg   *P2PConfig
	Peers []string
}

// NewHost creates a new host using the given configuration.
func NewHost(cfg *P2PConfig) (*Host, error) {
	return &Host{cfg: cfg, Peers: cfg.BootstrapPeers}, nil
}
