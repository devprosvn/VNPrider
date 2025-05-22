package crypto

import "testing"

func TestGenerateAddress(t *testing.T) {
	addr := GenerateAddress([]byte("pubkey"))
	if len(addr) == 0 {
		t.Fatalf("empty address")
	}
}
