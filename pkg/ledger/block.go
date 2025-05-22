// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package ledger

import "github.com/devprosvn/VNPrider/pkg/ledger/merkle"

// BlockHeader contains metadata about a block.
type BlockHeader struct {
	Height     uint64
	PrevHash   []byte
	MerkleRoot []byte
	Timestamp  int64
}

// Block is a group of transactions committed together.
type Block struct {
	Header       BlockHeader
	Transactions []Transaction
}

// ComputeMerkleRoot calculates the Merkle root for a set of transactions.
func ComputeMerkleRoot(txs []Transaction) []byte {
	var serialized [][]byte
	for _, tx := range txs {
		b, _ := tx.Serialize()
		serialized = append(serialized, b)
	}
	root, _ := merkle.BuildTree(serialized)
	return root
}
