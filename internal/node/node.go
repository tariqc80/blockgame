package node

import "github.com/tariqc80/blockgame/internal/block"
import "github.com/tariqc80/blockgame/internal/node/msg"

// Result interface for a return result
type Result interface{}

// NetworkNode interface for a Node that exists on a network
type NetworkNode interface {
	ID() string
	Socket() interface{}
	Broadcast(msg.Message) Result
}

// Peer struct represents a node in the blockchain network
// implements NetworkNode
type Peer struct {
	id       string
	chain    *block.Link
	messages *msg.Queue
}
