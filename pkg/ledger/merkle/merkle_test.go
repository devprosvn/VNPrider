// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package merkle

import (
	"testing"
)

func TestBuildTreeEmpty(t *testing.T) {
	_, err := BuildTree([][]byte{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildTreeTwo(t *testing.T) {
	txs := [][]byte{[]byte("a"), []byte("b")}
	_, err := BuildTree(txs)
	if err != nil {
		t.Fatal(err)
	}
}
