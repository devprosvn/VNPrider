// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package network

import "testing"

func TestNewHost(t *testing.T) {
	cfg := &P2PConfig{ListenPort: 0, BootstrapPeers: []string{}}
	NewHost(cfg)
}
