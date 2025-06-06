package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
	"github.com/devprosvn/VNPrider/pkg/node"
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
	if err := node.Run(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestRunNodeConfigError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := node.Run(ctx); err == nil {
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
	runNodeFn = node.Run
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

	old := node.NewLevelDBStore
	defer func() { node.NewLevelDBStore = old }()
	node.NewLevelDBStore = func(string) (*storage.LevelDBStore, error) { return nil, os.ErrNotExist }

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := node.Run(ctx); err == nil {
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

	oldHost := node.NewHost
	defer func() { node.NewHost = oldHost }()
	node.NewHost = func(*network.P2PConfig) (*network.Host, error) { return nil, os.ErrPermission }
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := node.Run(ctx); err == nil {
		t.Fatalf("expected host error")
	}
}

func TestRunNodeRPCPortInUse(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	port := ln.Addr().(*net.TCPAddr).Port

	os.WriteFile("config.toml", []byte(fmt.Sprintf("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=%d", port)), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()
	err = node.Run(ctx)
	ln.Close()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunNodeP2PInitFail(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\""), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	oldHost := node.NewHost
	defer func() { node.NewHost = oldHost }()
	node.NewHost = func(*network.P2PConfig) (*network.Host, error) { return nil, fmt.Errorf("p2p init fail") }
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := node.Run(ctx); err == nil || !strings.Contains(err.Error(), "p2p init fail") {
		t.Fatalf("expected p2p init fail error")
	}
}
