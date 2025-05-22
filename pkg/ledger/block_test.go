package ledger

import "testing"

func TestComputeMerkleRoot(t *testing.T) {
	tx := Transaction{Nonce: 1}
	root := ComputeMerkleRoot([]Transaction{tx})
	if len(root) == 0 {
		t.Fatalf("empty root")
	}
}
