package internal

import (
	"os"
	"testing"
)

func TestParseConfig(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)
	os.WriteFile("config.toml", []byte("data_dir=\"d\"\np2p.listen_port=1000\np2p.bootstrap_peers=[\"a\",\"b\"]\nrpc.listen_port=2000"), 0o644)
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	os.WriteFile("security.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\"\nrate_limit=5\nip_whitelist=[\"127.0.0.1\"]"), 0o644)

	cfg, err := ParseConfig()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DataDir != "d" || cfg.P2P.ListenPort != 1000 || cfg.RPC.ListenPort != 2000 {
		t.Fatalf("unexpected config %+v", cfg)
	}
	if len(cfg.P2P.BootstrapPeers) != 2 {
		t.Fatalf("bootstrap peers not parsed")
	}
	if cfg.Security.RateLimit != 5 || len(cfg.Security.IPWhitelist) != 1 {
		t.Fatalf("security not parsed")
	}
}

func TestParseConfigInvalidValidator(t *testing.T) {
	tmp := struct {
		Validators []ValidatorConfig `toml:"validator"`
	}{}
	os.WriteFile("bad.toml", []byte("validator[0].weight=abc"), 0o644)
	if err := parseFile("bad.toml", &tmp); err == nil {
		t.Fatalf("expected error")
	}
}

func TestParseFileSecurity(t *testing.T) {
	os.WriteFile("sec.toml", []byte("tls_cert_path=\"c\"\ntls_key_path=\"k\"\nrate_limit=3\nip_whitelist=[\"1.1.1.1\",\"2.2.2.2\"]"), 0o644)
	var s SecurityConfig
	if err := parseFile("sec.toml", &s); err != nil {
		t.Fatal(err)
	}
	if s.RateLimit != 3 || len(s.IPWhitelist) != 2 {
		t.Fatalf("security parse failed")
	}
}

func TestParseFileConfig(t *testing.T) {
	os.WriteFile("cfg.toml", []byte("data_dir=\"d\"\np2p.listen_port=1\np2p.bootstrap_peers=[\"x\"]\nrpc.listen_port=2"), 0o644)
	var c Config
	if err := parseFile("cfg.toml", &c); err != nil {
		t.Fatal(err)
	}
	if c.P2P.ListenPort != 1 || c.RPC.ListenPort != 2 || len(c.P2P.BootstrapPeers) != 1 {
		t.Fatalf("config parse failed")
	}
}
