package storage

import "github.com/devprosvn/VNPrider/pkg/ledger/block"

// BlockStore defines methods for block persistence.
type BlockStore interface {
	PutBlock(hash []byte, blk *block.Block) error
	GetBlock(hash []byte) (*block.Block, error)
	GetBlockByHeight(height uint64) (*block.Block, error)
	DeleteBlock(hash []byte) error
}

// StateStore defines methods for world state persistence.
type StateStore interface {
	Get(key []byte) ([]byte, error)
	Put(key, value []byte) error
	Delete(key []byte) error
}
