// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package mnemonic

import (
	"crypto/rand"
	"crypto/sha256"
	"strings"
)

// GenerateMnemonic generates a new mnemonic phrase
// GenerateMnemonic creates a mnemonic from random entropy.
func GenerateMnemonic() string {
	words := make([]string, 12)
	for i := range words {
		b := make([]byte, 1)
		rand.Read(b)
		idx := int(b[0]) % len(WordList)
		words[i] = WordList[idx]
	}
	return strings.Join(words, "-")
}

// ValidateMnemonic checks validity
// ValidateMnemonic checks whether all words exist in the list.
func ValidateMnemonic(m string) bool {
	for _, w := range strings.Split(m, "-") {
		found := false
		for _, v := range WordList {
			if w == v {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// MnemonicToSeed converts a mnemonic to a 32-byte seed.
func MnemonicToSeed(m string) []byte {
	sum := sha256.Sum256([]byte(m))
	return sum[:]
}
