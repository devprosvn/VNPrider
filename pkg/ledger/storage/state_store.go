// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package storage

import "github.com/devprosvn/VNPrider/pkg/ledger"

// BlockStore abstracts block storage.
type BlockStore interface {
	SaveBlock(block *ledger.Block) error
	GetBlock(height uint64) (*ledger.Block, error)
}

// StateStore abstracts state storage.
type StateStore interface {
	SetState(key string, value []byte) error
	GetState(key string) ([]byte, error)
}
