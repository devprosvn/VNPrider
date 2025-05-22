// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package crypto

import "fmt"

// GenerateAddress creates an address from the public key
// GenerateAddress derives an address string from the public key.
func GenerateAddress(pub []byte) string {
	sha := ComputeSHA3(pub)
	rip := ComputeRIPEMD160(sha)
	return fmt.Sprintf("%x", rip)
}
