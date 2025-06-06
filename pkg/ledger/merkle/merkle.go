// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package merkle

import "crypto/sha256"

// BuildTree constructs a Merkle root from transactions
// BuildTree constructs a Merkle root from transactions.
// BuildTree constructs a Merkle root from serialized transactions.
func BuildTree(txs [][]byte) ([]byte, error) {
	if len(txs) == 0 {
		h := sha256.Sum256(nil)
		return h[:], nil
	}
	var leaves [][]byte
	for _, b := range txs {
		sum := sha256.Sum256(b)
		leaves = append(leaves, sum[:])
	}
	for len(leaves) > 1 {
		var level [][]byte
		for i := 0; i < len(leaves); i += 2 {
			if i+1 == len(leaves) {
				level = append(level, leaves[i])
				continue
			}
			data := append(leaves[i], leaves[i+1]...)
			sum := sha256.Sum256(data)
			level = append(level, sum[:])
		}
		leaves = level
	}
	return leaves[0], nil
}
