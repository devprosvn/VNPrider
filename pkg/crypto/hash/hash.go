package hash

import (
	"golang.org/x/crypto/sha3"
)

// ComputeSHA3 returns the SHA3-256 digest of data.
func ComputeSHA3(data []byte) []byte {
	h := sha3.New256()
	h.Write(data)
	return h.Sum(nil)
}
