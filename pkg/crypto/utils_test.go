package crypto

import (
	"io"
	"testing"
)

func TestDoubleSHA256(t *testing.T) {
	h1 := DoubleSHA256([]byte("a"))
	h2 := DoubleSHA256([]byte("a"))
	if string(h1) != string(h2) {
		t.Fatalf("hash mismatch")
	}
}

func TestHMACSHA256(t *testing.T) {
	mac1 := HMACSHA256([]byte("k"), []byte("m"))
	mac2 := HMACSHA256([]byte("k"), []byte("m"))
	if string(mac1) != string(mac2) {
		t.Fatalf("mac mismatch")
	}
}

func TestRandomBytes(t *testing.T) {
	b, err := RandomBytes(5)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) != 5 {
		t.Fatalf("length wrong")
	}
}

func TestRandomBytesError(t *testing.T) {
	old := randReader
	randReader = errReader{}
	defer func() { randReader = old }()
	if _, err := RandomBytes(1); err == nil {
		t.Fatalf("expected error")
	}
}

type errReader struct{}

func (errReader) Read(p []byte) (int, error) { return 0, io.EOF }
