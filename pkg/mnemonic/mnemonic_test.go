package mnemonic

import "testing"

func TestMnemonicRoundtrip(t *testing.T) {
	m := GenerateMnemonic()
	if !ValidateMnemonic(m) {
		t.Fatalf("generated mnemonic invalid")
	}
	seed := MnemonicToSeed(m)
	if len(seed) != 32 {
		t.Fatalf("seed size incorrect")
	}
}

func TestValidateMnemonic(t *testing.T) {
	if ValidateMnemonic("bad-word") {
		t.Fatalf("expected invalid")
	}
}
