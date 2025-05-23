package integration

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/devprosvn/VNPrider/pkg/network"
	"github.com/devprosvn/VNPrider/pkg/node"
)

func freePort(t *testing.T) int {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func waitServer(t *testing.T, addr string) {
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		resp, err := http.Get(addr + "/status")
		if err == nil {
			resp.Body.Close()
			if resp.StatusCode == http.StatusOK {
				return
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Fatalf("server not ready: %s", addr)
}

func writeConfig(t *testing.T, dir string, port int) {
	cfg := fmt.Sprintf("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=%d", port)
	os.WriteFile(filepath.Join(dir, "config.toml"), []byte(cfg), 0o644)
	os.WriteFile(filepath.Join(dir, "validators.toml"), []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile(filepath.Join(dir, "security.toml"), []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)
}

func startNode(t *testing.T, dir string, ctx context.Context) chan error {
	old, _ := os.Getwd()
	os.Chdir(dir)
	errCh := make(chan error, 1)
	go func() {
		errCh <- node.Run(ctx)
	}()
	time.Sleep(50 * time.Millisecond)
	os.Chdir(old)
	return errCh
}

func TestNodeStartup_EndToEnd(t *testing.T) {
	if !Dummy() {
		t.Fatalf("dummy failed")
	}
	dir := t.TempDir()
	port := freePort(t)
	writeConfig(t, dir, port)
	ctx, cancel := context.WithCancel(context.Background())
	errCh := startNode(t, dir, ctx)
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port))
	cancel()
	if err := <-errCh; err != nil {
		t.Fatalf("node.Run error: %v", err)
	}
}

func TestGossipBetweenPeers(t *testing.T) {
	var mu sync.Mutex
	count := 0
	origNewHost := node.NewHost
	node.NewHost = func(cfg *network.P2PConfig) (*network.Host, error) {
		mu.Lock()
		count++
		mu.Unlock()
		return origNewHost(cfg)
	}
	defer func() { node.NewHost = origNewHost }()

	dir1 := t.TempDir()
	dir2 := t.TempDir()
	port1 := freePort(t)
	port2 := freePort(t)
	writeConfig(t, dir1, port1)
	writeConfig(t, dir2, port2)

	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	errCh1 := startNode(t, dir1, ctx1)
	errCh2 := startNode(t, dir2, ctx2)
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port1))
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port2))
	cancel1()
	cancel2()
	if err := <-errCh1; err != nil {
		t.Fatalf("node1 error: %v", err)
	}
	if err := <-errCh2; err != nil {
		t.Fatalf("node2 error: %v", err)
	}
	mu.Lock()
	defer mu.Unlock()
	if count < 2 {
		t.Fatalf("newHost not called for both nodes")
	}
}
