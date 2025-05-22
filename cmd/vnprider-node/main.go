// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/devprosvn/VNPrider/internal"
	"github.com/devprosvn/VNPrider/pkg/api"
	"github.com/devprosvn/VNPrider/pkg/consensus"
	"github.com/devprosvn/VNPrider/pkg/ledger"
	"github.com/devprosvn/VNPrider/pkg/ledger/storage"
	"github.com/devprosvn/VNPrider/pkg/network"
)

func main() {
	log.Println("vnprider-node starting...")
	cfg, err := internal.ParseConfig()
	internal.CheckErr(err)

	store, err := storage.NewLevelDBStore(cfg.DataDir)
	internal.CheckErr(err)
	_ = store

	host, err := network.NewHost(&network.P2PConfig{ListenPort: cfg.P2P.ListenPort, BootstrapPeers: cfg.P2P.BootstrapPeers})
	internal.CheckErr(err)
	gossip := network.NewGossip(host)

	engine := consensus.NewEngine([]consensus.Validator{}, func(b *ledger.Block) {
		data, _ := json.Marshal(b)
		gossip.BroadcastProposal(data)
	})

	apiSrv := api.NewServer()
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.RPC.ListenPort), apiSrv))
	}()

	engine.EnterNewRound()
	select {}
}
