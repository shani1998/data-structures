package queue

import (
	"errors"
)

// Queue implements thread-safe queue
// using buffered channel
type Queue struct {
	capacity int
	items    chan any
}

// NewQueue initialize new queue with capacity
func NewQueue(cap int) *Queue {
	return &Queue{
		capacity: cap,
		items:    make(chan any, cap),
	}
}

// Push enqueues an element. Returns error if queue is full.
func (q *Queue) Push(val any) error {
	if len(q.items) < q.capacity {
		q.items <- val
	}
	return errors.New("queue is full")
}

// Pop dequeues an element. Returns error if queue is empty.
func (q *Queue) Pop() (any, error) {
	if len(q.items) > 0 {
		item := <-q.items
		return item, nil
	}
	return nil, errors.New("queue is empty")
}

// Length is a wrapper for `len(s.items)`
func (q *Queue) Length() int {
	return len(q.items)
}
