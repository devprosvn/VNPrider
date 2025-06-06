// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package consensus

import "github.com/devprosvn/VNPrider/pkg/ledger"

// MsgType enumerates consensus message types

type MsgType int

const (
	MsgProposal MsgType = iota
	MsgPrevote
	MsgPrecommit
)

// Msg is a generic consensus message

type Msg struct {
	Type    MsgType
	Payload []byte
}

// Vote represents a validator vote.
type Vote struct {
	ValidatorID string
	Height      uint64
	Round       int
	Hash        []byte
}

// ProposalMsg carries a proposed block.
type ProposalMsg struct {
	Height uint64
	Round  int
	Block  *ledger.Block
}
