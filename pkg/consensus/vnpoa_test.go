package consensus

import (
	"encoding/json"
	"testing"

	"github.com/devprosvn/VNPrider/pkg/ledger"
)

func TestEngineProposeBlock(t *testing.T) {
	var got *ledger.Block
	e := NewEngine(nil, func(b *ledger.Block) { got = b })
	tx := ledger.Transaction{Nonce: 1}
	e.mempool = []ledger.Transaction{tx}
	e.ProposeBlock()
	if e.height != 1 {
		t.Fatalf("height not incremented")
	}
	if got == nil || got.Header.Height != 1 {
		t.Fatalf("broadcaster not called")
	}
	root := ledger.ComputeMerkleRoot(e.mempool)
	if string(root) != string(got.Header.MerkleRoot) {
		t.Fatalf("merkle root mismatch")
	}
}

func TestEngineHandleMessage(t *testing.T) {
	var got *ledger.Block
	e := NewEngine(nil, func(b *ledger.Block) { got = b })
	block := &ledger.Block{Header: ledger.BlockHeader{Height: 2}}
	p := ProposalMsg{Height: 2, Round: 0, Block: block}
	payload, _ := json.Marshal(p)
	e.HandleMessage(Msg{Type: MsgProposal, Payload: payload})
	if e.height != 2 {
		t.Fatalf("height not updated")
	}
	if got == nil || got.Header.Height != 2 {
		t.Fatalf("broadcaster not called")
	}
}

func TestEngineEnterNewRound(t *testing.T) {
	var got *ledger.Block
	e := NewEngine(nil, func(b *ledger.Block) { got = b })
	e.EnterNewRound()
	if got == nil {
		t.Fatalf("no block proposed")
	}
}
