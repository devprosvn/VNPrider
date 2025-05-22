package integration

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
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

func TestNodeEndToEnd(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)

	port := freePort(t)
	cfg := fmt.Sprintf("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=%d", port)
	os.WriteFile("config.toml", []byte(cfg), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)

	rootDir := filepath.Dir(filepath.Dir(old))
	bin := filepath.Join(dir, "node")
	build := exec.Command("go", "build", "-o", bin, "./cmd/vnprider-node")
	build.Dir = rootDir
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build error: %v\n%s", err, out)
	}

	cmd := exec.Command(bin)
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port))
}
