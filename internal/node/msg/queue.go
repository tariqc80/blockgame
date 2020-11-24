package msg

import "container/list"

// Queue struct is a message queue
type Queue struct {
	l *list.List
}

// NewQueue creates and returns a new message queue
func NewQueue() *Queue {
	q := new(Queue)
	q.l = list.New()

	return q
}

// Add adds a message to the queue
func (q *Queue) Add(m *Message) {
	q.l.PushBack(m)
}

// Remove pops the next message on the list and returns it
func (q *Queue) Remove() Message {
	return q.l.Remove(q.l.Back()).(Message)
}

// Read reads the next message in the queue without removing it
func (q *Queue) Read() Message {
	return q.l.Front().Value.(Message)
}
