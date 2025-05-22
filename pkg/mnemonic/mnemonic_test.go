// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package mnemonic

import "testing"

func TestMnemonicRoundtrip(t *testing.T) {
	m := GenerateMnemonic()
	if !ValidateMnemonic(m) {
		t.Fail()
	}
	_ = MnemonicToSeed(m)
}
