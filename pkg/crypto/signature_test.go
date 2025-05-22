package crypto

import (
	"crypto/rand"
	"testing"
)

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

func TestGenerateKeypairError(t *testing.T) {
	oldReader := rand.Reader
	rand.Reader = errReader{}
	defer func() { rand.Reader = oldReader }()
	if _, _, err := GenerateKeypair(); err == nil {
		t.Fatalf("expected error")
	}
}
