// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"github.com/devprosvn/VNPrider/pkg/ledger"
)

// LevelDBStore implements BlockStore and StateStore

type LevelDBStore struct {
	mu     sync.RWMutex
	blocks map[uint64]*ledger.Block
	state  map[string][]byte
}

func NewLevelDBStore(path string) (*LevelDBStore, error) {
	return &LevelDBStore{
		blocks: make(map[uint64]*ledger.Block),
		state:  make(map[string][]byte),
	}, nil
}

func (l *LevelDBStore) SaveBlock(b *ledger.Block) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.blocks[b.Header.Height] = b
	return nil
}

func (l *LevelDBStore) GetBlock(height uint64) (*ledger.Block, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	b, ok := l.blocks[height]
	if !ok {
		return nil, errors.New("block not found")
	}
	return b, nil
}

func (l *LevelDBStore) SetState(key string, value []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.state[key] = value
	return nil
}

func (l *LevelDBStore) GetState(key string) ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	v, ok := l.state[key]
	if !ok {
		return nil, errors.New("state not found")
	}
	return v, nil
}

// ExportSnapshot exports database snapshot
func (l *LevelDBStore) ExportSnapshot(path string) error {
	l.mu.RLock()
	defer l.mu.RUnlock()
	snap := struct {
		Blocks map[uint64]*ledger.Block
		State  map[string][]byte
	}{l.blocks, l.state}
	data, err := json.Marshal(snap)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// ImportSnapshot imports database snapshot
func (l *LevelDBStore) ImportSnapshot(path string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	snap := struct {
		Blocks map[uint64]*ledger.Block
		State  map[string][]byte
	}{}
	if err := json.Unmarshal(data, &snap); err != nil {
		return err
	}
	l.blocks = snap.Blocks
	l.state = snap.State
	return nil
}
