// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package crypto

import "testing"

func TestSignVerify(t *testing.T) {
	priv, pub, _ := GenerateKeypair()
	msg := []byte("hello")
	sig, _ := Sign(priv, msg)
	Verify(pub, msg, sig)
}
