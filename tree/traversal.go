package tree

import "fmt"

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%v ", root.val)
}

func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%v ", root.val)
	InOrderTraversal(root.right)
}
