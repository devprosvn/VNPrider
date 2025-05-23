package node

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devprosvn/VNPrider/internal"
	"github.com/devprosvn/VNPrider/pkg/api"
	"github.com/devprosvn/VNPrider/pkg/consensus"
	"github.com/devprosvn/VNPrider/pkg/ledger"
	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
)

var NewLevelDBStore = storage.NewLevelDBStore
var NewHost = network.NewHost

// Run starts the VNPrider node.
func Run(ctx context.Context) error {
	cfg, err := internal.ParseConfig()
	if err != nil {
		return err
	}
	store, err := NewLevelDBStore(cfg.DataDir)
	if err != nil {
		return err
	}
	_ = store
	host, err := NewHost(&network.P2PConfig{ListenPort: cfg.P2P.ListenPort, BootstrapPeers: cfg.P2P.BootstrapPeers})
	if err != nil {
		return err
	}
	gossip := network.NewGossip(host)
	engine := consensus.NewEngine([]consensus.Validator{}, func(b *ledger.Block) {
		data, _ := json.Marshal(b)
		gossip.BroadcastProposal(data)
	})
	apiSrv := api.NewServer()
	srv := &http.Server{Addr: fmt.Sprintf(":%d", cfg.RPC.ListenPort), Handler: apiSrv}
	go srv.ListenAndServe()
	engine.EnterNewRound()
	<-ctx.Done()
	return srv.Shutdown(context.Background())
}
