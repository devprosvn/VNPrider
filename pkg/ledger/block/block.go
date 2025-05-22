package block

import "github.com/devprosvn/VNPrider/pkg/ledger/types"

type BlockHeader struct {
	Version       uint32
	Height        uint64
	PrevBlockHash []byte
	MerkleRoot    []byte
	StateRoot     []byte
	Timestamp     int64
	ConsensusData []byte
}

type Block struct {
	Header       BlockHeader
	Transactions []types.Transaction
}
