package linkedlist

import (
	"fmt"
)

// SinglyLinkedList represents an actual linked list
// with its head/current reference.
type SinglyLinkedList struct {
	head *SinglyNode
	len  int
}

// SinglyNode represents a single node in the linked list.
type SinglyNode struct {
	Val  any
	Next *SinglyNode
}

// NewSinglyNode returns a new node with given Val.
func NewSinglyNode(Val interface{}) *SinglyNode {
	return &SinglyNode{
		Val:  Val,
		Next: nil,
	}
}

// InsertFirst inserts a node at the beginning of the linked list.
func (l *SinglyLinkedList) InsertFirst(node *SinglyNode) {
	// if there is no node in the list
	if l.head == nil {
		l.head = node
		l.len++
		return
	}
	// Make Next of new node as head
	node.Next = l.head
	// move the head to point to the new node
	l.head = node
	l.len++
}

// InsertLast inserts a node at the end of the linked list.
func (l *SinglyLinkedList) InsertLast(node *SinglyNode) {
	// if there is no node in the list
	if l.head == nil {
		l.head = node
		l.len++
		return
	}
	// find last node in the list
	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	l.len++
}

// DeleteFirst delete a first node from the linked list.
func (l *SinglyLinkedList) DeleteFirst() {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// if there is only one element in the list
	if l.head.Next == nil {
		l.head = nil
		l.len--
		return
	}
	// delete first node
	prev := l.head
	l.head = l.head.Next
	prev.Next = nil // break the link
	l.len--
}

// DeleteLast delete a last node from the linked list.
func (l *SinglyLinkedList) DeleteLast() {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// find 2nd last element and break the link
	curr, prv := l.head, l.head
	for curr.Next != nil {
		prv = curr
		curr = curr.Next
	}
	// if there is only one element in the list
	if prv == curr {
		l.head = nil
		l.len--
		return
	}

	prv.Next = nil
	l.len--
}

// DeleteNode delete the first occurrence of a node with given Val.
func (l *SinglyLinkedList) DeleteNode(key interface{}) {
	// if there is no node in the list
	if l.head == nil {
		return
	}
	// if head itself holds the key
	if l.head.Val == key {
		l.DeleteFirst()
		return
	}
	// find prv and Next elements to the given Val and break the link
	curr, prv := l.head, l.head
	for curr.Next != nil {
		if curr.Val == key {
			break
		}
		prv = curr
		curr = curr.Next
	}

	prv.Next = curr.Next
	curr.Next = nil
	l.len--
}

// Traverse iterates over the receiver linked list.
func (l *SinglyLinkedList) Traverse() {
	currNode := l.head
	for currNode.Next != nil {
		fmt.Printf("%v-->", currNode.Val)
		currNode = currNode.Next
	}
	fmt.Printf("%v\nTotal Elements: %v\n", currNode.Val, l.len)
}
