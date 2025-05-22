// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package consensus

import "testing"

func TestEngineEnterNewRound(t *testing.T) {
	validators := []Validator{{ID: "v1"}, {ID: "v2"}, {ID: "v3"}, {ID: "v4"}}
	e := NewEngine(validators, nil)
	e.EnterNewRound()
}
