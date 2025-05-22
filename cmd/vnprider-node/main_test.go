package main

import (
	"context"
	"os"
	"testing"
	"time"
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
