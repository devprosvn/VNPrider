// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package storage

import "testing"

func TestNewLevelDBStore(t *testing.T) {
	_, _ = NewLevelDBStore("/tmp/vnprider-test")
}
