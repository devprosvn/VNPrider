package ripemd160

import (
	"crypto/sha1"
	"hash"
)

// New returns a placeholder RIPEMD-160 implementation using SHA1.
// This is for offline builds and SHOULD NOT be used in production.
func New() hash.Hash {
	return sha1.New()
}
