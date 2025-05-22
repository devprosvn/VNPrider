package storage

import (
	"os"
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

func TestLevelDBStoreErrors(t *testing.T) {
	store, err := NewLevelDBStore("")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.GetBlock(99); err == nil {
		t.Fatalf("expected error for missing block")
	}
	if _, err := store.GetState("missing"); err == nil {
		t.Fatalf("expected error for missing state")
	}
	badPath := filepath.Join("non", "exist", "snap.json")
	if err := store.ExportSnapshot(badPath); err == nil {
		t.Fatalf("expected export error")
	}
	os.WriteFile(badPath, []byte("bad"), 0o644)
	if err := store.ImportSnapshot(badPath); err == nil {
		t.Fatalf("expected import error")
	}

	if err := store.ImportSnapshot("missing.json"); err == nil {
		t.Fatalf("expected read error")
	}

	oldMarshal := jsonMarshal
	jsonMarshal = func(any) ([]byte, error) { return nil, os.ErrInvalid }
	defer func() { jsonMarshal = oldMarshal }()
	if err := store.ExportSnapshot(filepath.Join(t.TempDir(), "snap")); err == nil {
		t.Fatalf("expected marshal error")
	}

	oldUnmarshal := jsonUnmarshal
	jsonUnmarshal = func([]byte, any) error { return os.ErrInvalid }
	defer func() { jsonUnmarshal = oldUnmarshal }()
	path := filepath.Join(t.TempDir(), "snap2")
	os.WriteFile(path, []byte("{}"), 0o644)
	if err := store.ImportSnapshot(path); err == nil {
		t.Fatalf("expected unmarshal error")
	}
}
