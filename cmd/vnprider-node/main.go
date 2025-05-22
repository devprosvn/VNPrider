// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/devprosvn/VNPrider/internal"
	"github.com/devprosvn/VNPrider/pkg/api"
	"github.com/devprosvn/VNPrider/pkg/consensus"
	"github.com/devprosvn/VNPrider/pkg/ledger"
	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
)

var newLevelDBStore = storage.NewLevelDBStore
var newHost = network.NewHost

func runNode(ctx context.Context) error {
	cfg, err := internal.ParseConfig()
	if err != nil {
		return err
	}

	store, err := newLevelDBStore(cfg.DataDir)
	if err != nil {
		return err
	}
	_ = store

	host, err := newHost(&network.P2PConfig{ListenPort: cfg.P2P.ListenPort, BootstrapPeers: cfg.P2P.BootstrapPeers})
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

func main() {
	log.Println("vnprider-node starting...")
	ctx := context.Background()
	if err := runNode(ctx); err != nil {
		if os.Getenv("TESTING") != "" {
			panic(err)
		}
		log.Fatal(err)
	}
}
