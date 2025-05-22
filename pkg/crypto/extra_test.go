package crypto

import "testing"

func TestXOR(t *testing.T) {
	a := []byte{1, 2, 3}
	b := []byte{3, 2, 1}
	x := XOR(a, b)
	expected := []byte{2, 0, 2}
	for i := range x {
		if x[i] != expected[i] {
			t.Fatalf("xor wrong")
		}
	}
}

func TestReverse(t *testing.T) {
	r := Reverse([]byte{1, 2, 3})
	if r[0] != 3 || r[2] != 1 {
		t.Fatalf("reverse failed")
	}
}

func TestChecksum32(t *testing.T) {
	c1 := Checksum32([]byte("a"))
	c2 := Checksum32([]byte("a"))
	if c1 != c2 {
		t.Fatalf("checksum inconsistent")
	}
}

func TestPadZero(t *testing.T) {
	out := PadZero([]byte{1, 2, 3}, 4)
	if len(out) != 4 || out[3] != 0 {
		t.Fatalf("pad failed")
	}
	out2 := PadZero([]byte{1, 2}, 0)
	if len(out2) != 2 {
		t.Fatalf("pad zero noop failed")
	}
	out3 := PadZero([]byte{1, 2}, 1)
	if len(out3) != 2 {
		t.Fatalf("pad already aligned")
	}
}
