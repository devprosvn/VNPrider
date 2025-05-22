// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package consensus

import (
	"encoding/json"
	"time"

	"github.com/devprosvn/VNPrider/pkg/ledger"
)

// Validator represents a consensus validator

type Validator struct {
	ID     string
	PubKey []byte
}

// Engine handles VNPOA consensus

type Engine struct {
	validators  []Validator
	height      uint64
	mempool     []ledger.Transaction
	broadcaster func(*ledger.Block)
}

// NewEngine creates a new consensus engine
// NewEngine creates a new consensus engine with the provided validators.
func NewEngine(validators []Validator, broadcaster func(*ledger.Block)) *Engine {
	return &Engine{validators: validators, broadcaster: broadcaster}
}

// EnterNewRound triggers the block proposal for the next height.
func (e *Engine) EnterNewRound() {
	e.ProposeBlock()
}

// HandleMessage processes incoming consensus messages.
func (e *Engine) HandleMessage(msg Msg) {
	switch msg.Type {
	case MsgProposal:
		var p ProposalMsg
		if err := json.Unmarshal(msg.Payload, &p); err != nil {
			return
		}
		if p.Block != nil && p.Block.Header.Height > e.height {
			e.height = p.Block.Header.Height
			if e.broadcaster != nil {
				e.broadcaster(p.Block)
			}
		}
	}
}

// ProposeBlock creates and broadcasts a new block proposal.
func (e *Engine) ProposeBlock() {
	header := ledger.BlockHeader{
		Height:    e.height + 1,
		Timestamp: time.Now().Unix(),
	}
	block := &ledger.Block{
		Header:       header,
		Transactions: e.mempool,
	}
	block.Header.MerkleRoot = ledger.ComputeMerkleRoot(block.Transactions)
	if e.broadcaster != nil {
		e.broadcaster(block)
	}
	e.height++
}
