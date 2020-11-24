package block

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/tariqc80/blockgame/pkg/util"
)

// Hash is an alias for the byte array
type Hash = string

// SummableData interface
type SummableData interface {
	sum() Hash
}

// Data interface
type Data struct {
	id           uint64 // unique id of block
	text         string // data the block holds
	previousHash Hash   // Checksum of the previous block's data
	timestamp    int64  // time the block was created. Universal for all peers

}

// NewData returns a new Data object
func NewData(t string, p Hash) *Data {
	return &Data{
		id:           rand.Uint64(),
		text:         t,
		previousHash: p,
		timestamp:    time.Now().Unix(),
	}
}

// Sum returns checksum of block data
func (d *Data) sum() Hash {
	h := sha256.New()
	h.Write(util.Uint64ToHexBytes(d.id))
	h.Write([]byte(d.text))
	h.Write([]byte(d.previousHash))
	h.Write(util.Int64ToHexBytes(d.timestamp))

	return hex.EncodeToString(h.Sum(nil))
}

// Timestamp returns the unix timestamp this block was created at
func (d *Data) Timestamp() int64 {
	return d.timestamp
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
