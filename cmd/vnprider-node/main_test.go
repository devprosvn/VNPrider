package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
)

func TestRunNode(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=0"), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	if err := runNode(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestRunNodeConfigError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := runNode(ctx); err == nil {
		t.Fatalf("expected error")
	}
}

func TestMainError(t *testing.T) {
	os.Setenv("TESTING", "1")
	defer os.Unsetenv("TESTING")
	defer func() { recover() }()
	main()
}

func TestMainFatal(t *testing.T) {
	os.Unsetenv("TESTING")
	called := false
	logFatal = func(...interface{}) { called = true }
	runNodeFn = func(context.Context) error { return os.ErrInvalid }
	main()
	if !called {
		t.Fatalf("logFatal not called")
	}
	runNodeFn = runNode
	logFatal = log.Fatal
}

func TestRunNodeStoreError(t *testing.T) {
	dir := t.TempDir()
	oldwd, _ := os.Getwd()
	defer os.Chdir(oldwd)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\""), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	old := newLevelDBStore
	defer func() { newLevelDBStore = old }()
	newLevelDBStore = func(string) (*storage.LevelDBStore, error) { return nil, os.ErrNotExist }

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := runNode(ctx); err == nil {
		t.Fatalf("expected store error")
	}
}

func TestRunNodeHostError(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\""), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	oldHost := newHost
	defer func() { newHost = oldHost }()
	newHost = func(*network.P2PConfig) (*network.Host, error) { return nil, os.ErrPermission }
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := runNode(ctx); err == nil {
		t.Fatalf("expected host error")
	}
}
