package node

import "github.com/tariqc80/blockgame/internal/block"

// NetworkNode interface for a Node that exists on a network
type NetworkNode interface {
	ID() string
	Socket() interface{}
	Broadcast(Message) MsgResult
}

// Peer struct represents a node in the blockchain network
// Implements NetworkNode
type Peer struct {
	id    string
	chain *block.Link
}
