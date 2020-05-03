package block

// Link interface
type Link interface {
	Previous() Link
	Validate(Link) bool
}

// Block struct
type Block struct {
	data     SummableData
	previous Link
}

// NewBlock creates a new Block instance
func NewBlock() Block {
	return Block{}
}

// Data returns the data struct for this block
func (b *Block) Data() SummableData {
	return b.data
}

// GetHash returns the unique hash for this block
func (b *Block) GetHash() Hash {
	return b.data.sum()
}

// Previous returns a pointer to the previous block
func (b *Block) Previous() Link {
	if b.previous != nil {
		return b.previous
	}
	return nil
}

// Validate chain
func (b *Block) Validate(o *Block) bool {
	if b.previous == nil && o.previous == nil {
		return true
	}

	if b.previous == nil || o.previous == nil {
		return false
	}

	if b.GetHash() != o.GetHash() {
		return false
	}

	return b.previous.Validate(o.previous)
}
