package crypto

import "testing"

func TestComputeSHA3(t *testing.T) {
	h := ComputeSHA3([]byte("test"))
	if len(h) != 32 {
		t.Fatalf("unexpected hash length")
	}
}

func TestComputeRIPEMD160(t *testing.T) {
	h := ComputeRIPEMD160([]byte("test"))
	if len(h) != 20 {
		t.Fatalf("unexpected hash length")
	}
}
