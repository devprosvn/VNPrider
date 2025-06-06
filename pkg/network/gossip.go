// Developed by DevPros with Codex as supporting tool
// DO NOT EDIT MANUALLY
package network

// Gossip handles network message broadcast

// Gossip handles broadcast and subscription of network messages.
type Gossip struct {
	subs chan []byte
}

// NewGossip constructs a Gossip instance for the provided host.
func NewGossip(host interface{}) *Gossip {
	return &Gossip{subs: make(chan []byte, 16)}
}

func (g *Gossip) BroadcastTx(tx []byte) {
	g.subs <- tx
}

func (g *Gossip) BroadcastProposal(header []byte) {
	g.subs <- header
}

func (g *Gossip) SubscribeMessages() <-chan []byte {
	return g.subs
}
