package msg

// MsgData inteface for the data of a queued message
type MsgData interface {
}

// Message struct represents a message
type Message struct {
	id        string
	data      MsgData
	timestamp int64
}
