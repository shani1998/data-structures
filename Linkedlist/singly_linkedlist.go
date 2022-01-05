package Linkedlist

import (
	"fmt"
)

// LinkedList represents an actual linked list
// with its head/current reference.
type LinkedList struct {
	head *Node
	len  int
}

// Node represents a single node in the linked list.
type Node struct {
	data interface{}
	next *Node
}

// NewNode returns a new node with given data.
func NewNode(data interface{}) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

// InsertFirst inserts a node at the beginning of the linked list.
func (l *LinkedList) InsertFirst(node *Node) {
	// if there is no node in the list
	if l.head == nil {
		l.head = node
		l.len++
		return
	}
	// insert given node  at the beginning of the list
	node.next = l.head
	l.head = node
	l.len++
}

// InsertLast inserts a node at the end of the linked list.
func (l *LinkedList) InsertLast(node *Node) {
	// if there is no node in the list
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
	curr.next = node
	l.len++
}

// DeleteFirst delete a first node from the linked list.
func (l *LinkedList) DeleteFirst() {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// if there is only one element in the list
	if l.head.next == nil {
		l.head = nil
		l.len--
		return
	}
	// delete first node
	prev := l.head
	l.head = l.head.next
	prev.next = nil // break the link
	l.len--
}

// DeleteLast delete a last node from the linked list.
func (l *LinkedList) DeleteLast() {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// find 2nd last element and break the link
	curr, prv := l.head, l.head
	for curr.next != nil {
		prv = curr
		curr = curr.next
	}
	// if there is only one element in the list
	if prv == curr {
		l.head = nil
		l.len--
		return
	}

	prv.next = nil
	l.len--
}

// DeleteNode delete the first occurrence of a node with given data.
func (l *LinkedList) DeleteNode(key interface{}) {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// if head itself holds the key
	if l.head.data == key {
		l.DeleteFirst()
		return
	}
	// find prv and next elements to the given data and break the link
	curr, prv := l.head, l.head
	for curr.next != nil {
		if curr.data == key {
			break
		}
		prv = curr
		curr = curr.next
	}

	prv.next = curr.next
	curr.next = nil
	l.len--
}

// Traverse iterates over the receiver linked list.
func (l *LinkedList) Traverse() {
	currNode := l.head
	for currNode.next != nil {
		fmt.Printf("%v-->", currNode.data)
		currNode = currNode.next
	}
	fmt.Printf("%v\nTotal Elements: %v\n", currNode.data, l.len)
}
