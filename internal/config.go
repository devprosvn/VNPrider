// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package internal

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// P2PConfig holds peer-to-peer settings.
type P2PConfig struct {
	ListenPort     int      `toml:"listen_port"`
	BootstrapPeers []string `toml:"bootstrap_peers"`
}

// RPCConfig holds RPC server configuration.
type RPCConfig struct {
	ListenPort int `toml:"listen_port"`
}

// ValidatorConfig describes a validator.
type ValidatorConfig struct {
	ID       string `toml:"id"`
	PubKey   string `toml:"pubkey"`
	Endpoint string `toml:"endpoint"`
	Weight   int    `toml:"weight"`
}

// SecurityConfig holds security related options.
type SecurityConfig struct {
	TLSCert     string   `toml:"tls_cert_path"`
	TLSKey      string   `toml:"tls_key_path"`
	RateLimit   int      `toml:"rate_limit"`
	IPWhitelist []string `toml:"ip_whitelist"`
}

// Config aggregates all node configuration.
type Config struct {
	DataDir    string    `toml:"data_dir"`
	P2P        P2PConfig `toml:"p2p"`
	RPC        RPCConfig `toml:"rpc"`
	Validators []ValidatorConfig
	Security   SecurityConfig
}

// ParseConfig loads configuration from default files.
func ParseConfig() (*Config, error) {
	cfg := &Config{}
	if err := parseFile("config.toml", cfg); err != nil {
		return nil, err
	}
	validators := struct {
		Validators []ValidatorConfig `toml:"validator"`
	}{}
	if err := parseFile("validators.toml", &validators); err != nil {
		return nil, err
	}
	cfg.Validators = validators.Validators
	if err := parseFile("security.toml", &cfg.Security); err != nil {
		return nil, err
	}
	return cfg, nil
}

func parseFile(path string, v any) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	m := make(map[string]string)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.Trim(strings.TrimSpace(parts[1]), "\"")
		m[key] = val
	}
	switch c := v.(type) {
	case *Config:
		if v, ok := m["data_dir"]; ok {
			c.DataDir = v
		}
		if v, ok := m["p2p.listen_port"]; ok {
			p, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("invalid p2p.listen_port: %w", err)
			}
			c.P2P.ListenPort = p
		}
		if v, ok := m["p2p.bootstrap_peers"]; ok {
			v = strings.TrimPrefix(strings.TrimSuffix(v, "]"), "[")
			if v != "" {
				for _, s := range strings.Split(v, ",") {
					s = strings.Trim(strings.TrimSpace(s), "\"")
					if s != "" {
						c.P2P.BootstrapPeers = append(c.P2P.BootstrapPeers, s)
					}
				}
			}
		}
		if v, ok := m["rpc.listen_port"]; ok {
			p, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("invalid rpc.listen_port: %w", err)
			}
			c.RPC.ListenPort = p
		}
	case *SecurityConfig:
		if v, ok := m["tls_cert_path"]; ok {
			c.TLSCert = v
		}
		if v, ok := m["tls_key_path"]; ok {
			c.TLSKey = v
		}
		if v, ok := m["rate_limit"]; ok {
			i, err := strconv.Atoi(v)
			if err != nil {
				return fmt.Errorf("invalid rate_limit: %w", err)
			}
			c.RateLimit = i
		}
		if v, ok := m["ip_whitelist"]; ok {
			v = strings.TrimPrefix(strings.TrimSuffix(v, "]"), "[")
			if v != "" {
				for _, s := range strings.Split(v, ",") {
					s = strings.Trim(strings.TrimSpace(s), "\"")
					if s != "" {
						c.IPWhitelist = append(c.IPWhitelist, s)
					}
				}
			}
		}
	case *struct {
		Validators []ValidatorConfig `toml:"validator"`
	}:
		tmp := make(map[int]*ValidatorConfig)
		for k, vStr := range m {
			if !strings.HasPrefix(k, "validator") {
				continue
			}
			remainder := strings.TrimPrefix(k, "validator")
			idx := 0
			if strings.HasPrefix(remainder, "[") {
				end := strings.Index(remainder, "]")
				if end == -1 {
					return fmt.Errorf("invalid validator key %q", k)
				}
				idStr := remainder[1:end]
				i, err := strconv.Atoi(idStr)
				if err != nil {
					return fmt.Errorf("invalid validator index in %q: %w", k, err)
				}
				idx = i
				remainder = remainder[end+1:]
			}
			if strings.HasPrefix(remainder, ".") {
				remainder = remainder[1:]
			}
			cfg, ok := tmp[idx]
			if !ok {
				cfg = &ValidatorConfig{}
				tmp[idx] = cfg
			}
			switch remainder {
			case "id":
				cfg.ID = vStr
			case "pubkey":
				cfg.PubKey = vStr
			case "endpoint":
				cfg.Endpoint = vStr
			case "weight":
				w, err := strconv.Atoi(vStr)
				if err != nil {
					return fmt.Errorf("invalid weight for validator %d: %w", idx, err)
				}
				cfg.Weight = w
			}
		}
		indices := make([]int, 0, len(tmp))
		for i := range tmp {
			indices = append(indices, i)
		}
		sort.Ints(indices)
		for _, i := range indices {
			vc := tmp[i]
			if vc.ID == "" || vc.PubKey == "" || vc.Endpoint == "" {
				return fmt.Errorf("missing fields for validator %d", i)
			}
			c.Validators = append(c.Validators, *vc)
		}
	}
	return scanner.Err()
}
