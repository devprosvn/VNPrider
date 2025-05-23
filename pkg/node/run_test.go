package node

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
)

func TestRun(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=0"), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	ctx, cancel := context.WithCancel(context.Background())
	go func() { time.Sleep(10 * time.Millisecond); cancel() }()
	if err := Run(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestRunErrors(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\""), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	oldStore := NewLevelDBStore
	NewLevelDBStore = func(string) (*storage.LevelDBStore, error) { return nil, os.ErrNotExist }
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Run(ctx); err == nil {
		t.Fatalf("expected store error")
	}
	NewLevelDBStore = oldStore

	oldHost := NewHost
	NewHost = func(*network.P2PConfig) (*network.Host, error) { return nil, os.ErrPermission }
	ctx, cancel = context.WithCancel(context.Background())
	cancel()
	if err := Run(ctx); err == nil {
		t.Fatalf("expected host error")
	}
	NewHost = oldHost
}

func TestRunParseError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Run(ctx); err == nil {
		t.Fatalf("expected parse error")
	}
}
