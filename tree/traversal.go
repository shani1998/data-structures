package tree

import (
	"fmt"
)

// PreOrderTraversal time O(n), space:(n)
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

// PostOrderTraversal time O(n), space:(n)
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.Left)
	PostOrderTraversal(root.Right)
	fmt.Printf("%v ", root.Val)
}

// InOrderTraversal time O(n), space:(n)
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left)
	fmt.Printf("%v ", root.Val)
	InOrderTraversal(root.Right)
}
