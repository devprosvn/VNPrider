package network

import "testing"

func TestGossipBroadcast(t *testing.T) {
	g := NewGossip(nil)
	ch := g.SubscribeMessages()
	g.BroadcastTx([]byte("tx"))
	if string(<-ch) != "tx" {
		t.Fatalf("tx not received")
	}
	g.BroadcastProposal([]byte("p"))
	if string(<-ch) != "p" {
		t.Fatalf("proposal not received")
	}
}
