package stack

import (
	"errors"
	"sync"
)

// Stack is a simple LIFO stack for any type.
type Stack []any

var mu sync.Mutex

// Push adds an item to the top of the stack
func (s *Stack) Push(item any) {
	mu.Lock()
	defer mu.Unlock()
	*s = append(*s, item)
}

// Pop removes the top item from the stack and returns it. An error is returned
// if called on an empty stack.
func (s *Stack) Pop() (any, error) {
	length := s.Length()

	if s.IsEmpty() {
		return nil, errors.New("cannot pop from empty stack")
	}

	mu.Lock()
	defer mu.Unlock()
	item := (*s)[length-1]
	*s = (*s)[:length-1]

	return item, nil
}

// Length is a wrapper for `len(s)`
func (s *Stack) Length() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty, and false otherwise
func (s *Stack) IsEmpty() bool {
	return s.Length() == 0
}
