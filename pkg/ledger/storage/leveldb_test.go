package storage

import (
	"path/filepath"
	"testing"

	"github.com/devprosvn/VNPrider/pkg/ledger"
)

func TestLevelDBStore(t *testing.T) {
	store, err := NewLevelDBStore("")
	if err != nil {
		t.Fatal(err)
	}
	blk := &ledger.Block{Header: ledger.BlockHeader{Height: 1}}
	if err := store.SaveBlock(blk); err != nil {
		t.Fatal(err)
	}
	b, err := store.GetBlock(1)
	if err != nil || b.Header.Height != 1 {
		t.Fatalf("block retrieval failed")
	}
	if err := store.SetState("k", []byte("v")); err != nil {
		t.Fatal(err)
	}
	v, err := store.GetState("k")
	if err != nil || string(v) != "v" {
		t.Fatalf("state retrieval failed")
	}
	snap := filepath.Join(t.TempDir(), "snap.json")
	if err := store.ExportSnapshot(snap); err != nil {
		t.Fatal(err)
	}
	store2, _ := NewLevelDBStore("")
	if err := store2.ImportSnapshot(snap); err != nil {
		t.Fatal(err)
	}
	if _, err := store2.GetBlock(1); err != nil {
		t.Fatalf("snapshot block missing")
	}

	if _, err := store2.GetBlock(2); err == nil {
		t.Fatalf("expected missing block error")
	}
}
