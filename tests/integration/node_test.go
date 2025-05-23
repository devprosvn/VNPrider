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

// buildNode compiles the vnprider-node binary into the given directory.
func buildNode(t *testing.T, dir string) string {
	rootDir := filepath.Dir(filepath.Dir(getwd(t)))
	bin := filepath.Join(dir, "node")
	cmd := exec.Command("go", "build", "-o", bin, "./cmd/vnprider-node")
	cmd.Dir = rootDir
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build error: %v\n%s", err, out)
	}
	return bin
}

func getwd(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return wd
}

func writeConfig(t *testing.T, dir string, port int) {
	cfg := fmt.Sprintf("data_dir=\"d\"\np2p.listen_port=0\nrpc.listen_port=%d", port)
	os.WriteFile(filepath.Join(dir, "config.toml"), []byte(cfg), 0o644)
	os.WriteFile(filepath.Join(dir, "validators.toml"), []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile(filepath.Join(dir, "security.toml"), []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\""), 0o644)
}

func startNode(t *testing.T, bin, dir string) *exec.Cmd {
	cmd := exec.Command(bin)
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	return cmd
}

func TestNodeStartup(t *testing.T) {
	dir := t.TempDir()
	port := freePort(t)
	writeConfig(t, dir, port)
	bin := buildNode(t, dir)
	cmd := startNode(t, bin, dir)
	defer cmd.Process.Kill()
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port))
}

func TestTwoNodes(t *testing.T) {
	dir1 := t.TempDir()
	dir2 := t.TempDir()
	port1 := freePort(t)
	port2 := freePort(t)
	writeConfig(t, dir1, port1)
	writeConfig(t, dir2, port2)
	bin1 := buildNode(t, dir1)
	bin2 := buildNode(t, dir2)
	cmd1 := startNode(t, bin1, dir1)
	defer cmd1.Process.Kill()
	cmd2 := startNode(t, bin2, dir2)
	defer cmd2.Process.Kill()
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port1))
	waitServer(t, fmt.Sprintf("http://127.0.0.1:%d", port2))
}
