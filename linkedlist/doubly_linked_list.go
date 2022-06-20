package linkedlist

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
