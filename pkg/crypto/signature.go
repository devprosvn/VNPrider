// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
)

// GenerateKeypair returns a new Ed25519 private/public key pair.
func GenerateKeypair() ([]byte, []byte, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return priv, pub, nil
}

func Sign(priv []byte, msg []byte) ([]byte, error) {
	signature := ed25519.Sign(ed25519.PrivateKey(priv), msg)
	return signature, nil
}

func Verify(pub []byte, msg []byte, sig []byte) bool {
	return ed25519.Verify(ed25519.PublicKey(pub), msg, sig)
}
