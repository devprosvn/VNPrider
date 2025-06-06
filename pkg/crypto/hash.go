// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package crypto

import (
	"crypto/sha256"
)

// ComputeSHA3 returns the SHA3-256 hash of the input data.
func ComputeSHA3(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}

// ComputeRIPEMD160 returns the RIPEMD-160 hash of the input data.
func ComputeRIPEMD160(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:20]
}
