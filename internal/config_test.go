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

func TestParseFileErrors(t *testing.T) {
	if err := parseFile("nofile.toml", &Config{}); err == nil {
		t.Fatalf("expected error for missing file")
	}

	os.WriteFile("badport.toml", []byte("p2p.listen_port=abc"), 0o644)
	if err := parseFile("badport.toml", &Config{}); err == nil {
		t.Fatalf("expected port error")
	}

	os.WriteFile("valmissing.toml", []byte("validator[0].id=\"id1\""), 0o644)
	tmp := struct {
		Validators []ValidatorConfig `toml:"validator"`
	}{}
	if err := parseFile("valmissing.toml", &tmp); err == nil {
		t.Fatalf("expected validator field error")
	}

	os.WriteFile("rpcbad.toml", []byte("rpc.listen_port=bad"), 0o644)
	if err := parseFile("rpcbad.toml", &Config{}); err == nil {
		t.Fatalf("expected rpc port error")
	}

	os.WriteFile("ratebad.toml", []byte("rate_limit=bad"), 0o644)
	if err := parseFile("ratebad.toml", &SecurityConfig{}); err == nil {
		t.Fatalf("expected rate limit error")
	}

	os.WriteFile("valindexbad.toml", []byte("validator[abc].id=\"x\""), 0o644)
	if err := parseFile("valindexbad.toml", &tmp); err == nil {
		t.Fatalf("expected validator index error")
	}

	os.WriteFile("valkeybad.toml", []byte("validator[0.id=\"x\""), 0o644)
	if err := parseFile("valkeybad.toml", &tmp); err == nil {
		t.Fatalf("expected validator key error")
	}
}

func TestParseValidators(t *testing.T) {
	os.WriteFile("vals.toml", []byte("validator[0].id=\"id1\"\nvalidator[0].pubkey=\"pk1\"\nvalidator[0].endpoint=\"ep1\"\nvalidator[0].weight=2"), 0o644)
	tmp := struct {
		Validators []ValidatorConfig `toml:"validator"`
	}{}
	if err := parseFile("vals.toml", &tmp); err != nil {
		t.Fatal(err)
	}
	if len(tmp.Validators) != 1 || tmp.Validators[0].Weight != 2 {
		t.Fatalf("validator not parsed")
	}
}

func TestParseConfigErrors(t *testing.T) {
	dir := t.TempDir()
	old, _ := os.Getwd()
	defer os.Chdir(old)
	os.Chdir(dir)

	// missing config.toml
	if _, err := ParseConfig(); err == nil {
		t.Fatalf("expected error for missing config")
	}

	// config exists but validators missing
	os.WriteFile("config.toml", []byte("data_dir=\"d\""), 0o644)
	if _, err := ParseConfig(); err == nil {
		t.Fatalf("expected validator error")
	}

	// validators exist but security missing
	os.WriteFile("validators.toml", []byte("[validator]\nid=\"id1\"\npubkey=\"pk\"\nendpoint=\"ep\"\nweight=1"), 0o644)
	if _, err := ParseConfig(); err == nil {
		t.Fatalf("expected security error")
	}
}
