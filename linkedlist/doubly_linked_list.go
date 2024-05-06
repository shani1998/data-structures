package linkedlist

import (
	"errors"
	"fmt"
)

// LinkedList represents an actual doubly linked list
// with its head reference.
type LinkedList struct {
	head *Node
	len  int
}

// Node represents a single node in the doubly linked list.
type Node struct {
	data       interface{}
	next, prev *Node
}

func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
		prev: nil,
	}
}

// InsertFirst inserts a node at the beginning of the doubly linked list.
func (l *LinkedList) InsertFirst(node *Node) {
	// if the Linked List is empty, then
	// make the new node as head
	if l.head == nil {
		l.head = node
		l.len++
		return
	}
	// change prev of head node to the new node
	l.head.prev = node
	// Make next of new node as head
	node.next = l.head
	// move the head to point to the new node
	l.head = node
	l.len++
}

// InsertLast inserts a node at the end of the doubly linked list.
func (l *LinkedList) InsertLast(node *Node) {
	// if the Linked List is empty, then
	// make the new node as head
	if l.head == nil {
		l.head = node
		l.len++
		return
	}
	// find last node in the list
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	// Change the next of last nod
	curr.next = node
	// Make last node as previous of new node
	node.prev = curr
	l.len++
}

func (l *LinkedList) InsertAtIndex(index int, node *Node) error {
	if index < 0 || index > l.len {
		return errors.New("index out of bounds")
	}

	if index == 0 {
		l.InsertFirst(node)
		return nil
	}

	if index == l.len {
		l.InsertLast(node)
		return nil
	}

	i, curr, prev := 0, l.head, l.head
	for i != index {
		prev = curr
		curr = curr.next
		i++
	}

	node.prev = prev
	node.next = curr

	prev.next = node
	curr.prev = node

	l.len++

	return nil
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) TraverseForward() {
	curr := l.head
	if curr == nil {
		return
	}

	output := ""
	for curr != nil {
		output = fmt.Sprintf("%s-->[%v]", output, curr.data)
		curr = curr.next
	}

	fmt.Printf("%s\n", output)
}

func (l *LinkedList) TraverseBackward() {
	curr := l.head
	if curr == nil {
		return
	}
	// go to the last node
	for curr.next != nil {
		curr = curr.next
	}

	output := ""
	// traverse backward using prev pointer
	for curr != nil {
		output = fmt.Sprintf("%s-->[%v]", output, curr.data)
		curr = curr.prev
	}

	fmt.Printf("%s\n", output)
}
