package tree

import (
	"fmt"
)

// PreOrderTraversal time O(n), space:(n)
func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

// PostOrderTraversal time O(n), space:(n)
func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%v ", root.val)
}

// InOrderTraversal time O(n), space:(n)
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%v ", root.val)
	InOrderTraversal(root.right)
}
