package tree

import (
	"sync"

	"github.com/shani1998/data-structures/queue"
)

// Node a single node that composes the tree
type Node struct {
	Val   any
	Left  *Node
	Right *Node
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
		Val:   val,
		Left:  nil,
		Right: nil,
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
		if currentNode.(*Node).Left == nil {
			currentNode.(*Node).Left = node
			return
		}
		_ = q.Push(currentNode.(*Node).Left)

		if currentNode.(*Node).Right == nil {
			currentNode.(*Node).Right = node
			return
		}
		_ = q.Push(currentNode.(*Node).Right)
	}

}
