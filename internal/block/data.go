package block

import (
	"crypto/sha256"
	"time"
)

// Hash type for the hash value
type Hash [sha256.Size]byte

// SummableData interface
type SummableData interface {
	sum() Hash
}

// Data interface
type Data struct {
	text        string
	createdTime time.Time
}

// Sum returns checksum of block data
func (d *Data) Sum() Hash {
	// TODO add CreatedTime to sum
	return sha256.Sum256([]byte(d.text))
}

// CreatedTime returns the unix timestamp this block was created at
func (d *Data) CreatedTime() int64 {
	return d.createdTime.Unix()
}
