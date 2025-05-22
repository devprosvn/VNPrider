package merkle

import (
	"errors"

	"github.com/devprosvn/VNPrider/pkg/crypto/hash"
	"github.com/devprosvn/VNPrider/pkg/ledger/types"
)

// MerkleProof holds sibling hash and direction information.
type MerkleProof struct {
	Hash []byte
	Left bool
}

// BuildTree computes the Merkle root from transactions.
func BuildTree(txs []types.Transaction) ([]byte, error) {
	if len(txs) == 0 {
		return nil, errors.New("no transactions")
	}
	var hashes [][]byte
	for _, tx := range txs {
		hashes = append(hashes, hash.ComputeSHA3(tx.Serialize()))
	}
	for len(hashes) > 1 {
		var next [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				next = append(next, hashes[i])
			} else {
				combined := append(hashes[i], hashes[i+1]...)
				next = append(next, hash.ComputeSHA3(combined))
			}
		}
		hashes = next
	}
	return hashes[0], nil
}

// GenerateProof generates a Merkle proof for the given transaction ID.
func GenerateProof(txID []byte, txs []types.Transaction) ([]MerkleProof, error) {
	// simplified placeholder: not implemented
	return nil, errors.New("not implemented")
}
