package tree

import (
	"github.com/shani1998/data-structures/queue"
	"sync"
)

// Node a single node that composes the tree
type Node struct {
	val   any
	left  *Node
	right *Node
}

// BinaryTree the binary tree of values of type any
type BinaryTree struct {
	root *Node
	lock sync.RWMutex
}

// NewNode returns a new binary tree
// node with given val.
func NewNode(val any) *Node {
	return &Node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

// Insert inserts and node into the tree in level order manner.
func (bt *BinaryTree) Insert(node *Node) {
	bt.lock.Lock()
	defer bt.lock.Unlock()

	if bt.root == nil {
		bt.root = node
		return
	}

	// perform level order traversal until we find
	// an empty space.
	q := queue.NewQueue(100)
	_ = q.Push(bt.root)
	for q.Length() > 0 {
		currentNode, _ := q.Pop()
		if currentNode.(*Node).left == nil {
			currentNode.(*Node).left = node
			return
		}
		_ = q.Push(currentNode.(*Node).left)

		if currentNode.(*Node).right == nil {
			currentNode.(*Node).right = node
			return
		}
		_ = q.Push(currentNode.(*Node).right)
	}

}
