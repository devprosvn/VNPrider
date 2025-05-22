package crypto

import "testing"

func TestSignVerify(t *testing.T) {
	priv, pub, err := GenerateKeypair()
	if err != nil {
		t.Fatal(err)
	}
	msg := []byte("hello")
	sig, err := Sign(priv, msg)
	if err != nil {
		t.Fatal(err)
	}
	if !Verify(pub, msg, sig) {
		t.Fatalf("verify failed")
	}
}
