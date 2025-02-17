package bst

import "github.com/shani1998/data-structures/tree"

type BST struct {
	root *tree.Node
}

func (bst *BST) Insert(Val int) {
	bst.insertHelper(bst.root, Val)
}

func (bst *BST) insertHelper(node *tree.Node, Val int) *tree.Node {
	if bst.root == nil {
		bst.root = &tree.Node{Val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &tree.Node{Val, nil, nil}
	}

	if Val <= node.Val.(int) {
		node.Left = bst.insertHelper(node.Left, Val)
	}

	if Val > node.Val.(int) {
		node.Right = bst.insertHelper(node.Right, Val)
	}

	return node
}
