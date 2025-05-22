package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/devprosvn/VNPrider/pkg/mnemonic"
)

func runCLI(args ...string) string {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = append([]string{"vnprider"}, args...)
	r, w, _ := os.Pipe()
	oldOut := os.Stdout
	os.Stdout = w
	main()
	w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	return strings.TrimSpace(string(out))
}

func TestCLIMnemonic(t *testing.T) {
	out := runCLI("mnemonic")
	if !mnemonic.ValidateMnemonic(out) {
		t.Fatalf("invalid mnemonic %q", out)
	}
}

func TestCLIStatus(t *testing.T) {
	out := runCLI("status")
	if out != "node running" {
		t.Fatalf("unexpected output %q", out)
	}
}

func TestCLIUnknown(t *testing.T) {
	out := runCLI("badcmd")
	if out != "unknown command" {
		t.Fatalf("unexpected output %q", out)
	}
}

func TestCLIValidateUsage(t *testing.T) {
	out := runCLI("validate")
	if !strings.HasPrefix(out, "usage") {
		t.Fatalf("expected usage, got %q", out)
	}
}

func TestCLIValidate(t *testing.T) {
	m := mnemonic.GenerateMnemonic()
	out := runCLI("validate", m)
	if out != "valid" {
		t.Fatalf("expected valid got %q", out)
	}
}

func TestCLIValidateInvalid(t *testing.T) {
	out := runCLI("validate", "bad-word")
	if out != "invalid" {
		t.Fatalf("expected invalid got %q", out)
	}
}

func TestCLIUsage(t *testing.T) {
	out := runCLI()
	if !strings.HasPrefix(out, "usage:") {
		t.Fatalf("expected usage output, got %q", out)
	}
}
