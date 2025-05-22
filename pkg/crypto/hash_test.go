// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package crypto

import "testing"

func TestComputeSHA3(t *testing.T) {
	_ = ComputeSHA3([]byte("test"))
}

func TestComputeRIPEMD160(t *testing.T) {
	_ = ComputeRIPEMD160([]byte("test"))
}
