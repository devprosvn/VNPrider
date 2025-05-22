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

func TestBuildTreeSingle(t *testing.T) {
	txs := [][]byte{[]byte("a")}
	root, err := BuildTree(txs)
	if err != nil {
		t.Fatal(err)
	}
	if len(root) == 0 {
		t.Fatalf("empty root")
	}
}

func TestBuildTreeOdd(t *testing.T) {
	txs := [][]byte{[]byte("a"), []byte("b"), []byte("c")}
	root, err := BuildTree(txs)
	if err != nil {
		t.Fatal(err)
	}
	if len(root) == 0 {
		t.Fatalf("empty root")
	}
}
