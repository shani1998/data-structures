package Linkedlist

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
	data interface{}
	next *SinglyNode
}

// NewSinglyNode returns a new node with given data.
func NewSinglyNode(data interface{}) *SinglyNode {
	return &SinglyNode{
		data: data,
		next: nil,
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
	// Make next of new node as head
	node.next = l.head
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
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = node
	l.len++
}

// DeleteFirst delete a first node from the linked list.
func (l *SinglyLinkedList) DeleteFirst() {
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
func (l *SinglyLinkedList) DeleteLast() {
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
func (l *SinglyLinkedList) DeleteNode(key interface{}) {
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

// ReverseList reverse the Linked-list and return new head.
// i/p: 2-->3-->4-->5-->6, o/p: 6-->5-->4-->3-->2
func (l *SinglyLinkedList) ReverseList() *SinglyNode {
	// if there is no node in the list
	if l.head == nil {
		return nil
	}
	var prv *SinglyNode
	prv, curr := nil, l.head
	for curr != nil {
		temp := curr.next
		curr.next = prv
		prv = curr
		curr = temp
	}
	return prv
}

// Traverse iterates over the receiver linked list.
func (l *SinglyLinkedList) Traverse() {
	currNode := l.head
	for currNode.next != nil {
		fmt.Printf("%v-->", currNode.data)
		currNode = currNode.next
	}
	fmt.Printf("%v\nTotal Elements: %v\n", currNode.data, l.len)
}
