package signature

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateKeypair creates a new Ed25519 keypair.
func GenerateKeypair() (ed25519.PrivateKey, ed25519.PublicKey, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	return priv, pub, err
}

// Sign signs msg with priv.
func Sign(priv ed25519.PrivateKey, msg []byte) []byte {
	return ed25519.Sign(priv, msg)
}

// Verify verifies sig of msg under pub.
func Verify(pub ed25519.PublicKey, msg, sig []byte) bool {
	return ed25519.Verify(pub, msg, sig)
}
