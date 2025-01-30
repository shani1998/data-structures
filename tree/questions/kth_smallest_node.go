package questions

import (
	"fmt"

	"github.com/shani1998/data-structures/tree"
)

func kthSmallest(root *tree.Node, k int) int {
	count := 0
	var result int

	var inorder func(*tree.Node)
	inorder = func(node *tree.Node) {
		if node == nil || count >= k {
			return
		}

		inorder(node.Left)

		count++ // if k is allowed to modify, then we can dec k until it becomes 0
		if count == k {
			result = node.Val.(int)
			return
		}

		inorder(node.Right)
	}

	inorder(root)
	return result
}

func kthSmallest1(root *tree.Node, k int) int {
	var inorder func(*tree.Node, []int) []int
	inorder = func(root *tree.Node, inroderNodes []int) []int {
		if root == nil {
			return inroderNodes
		}
		inroderNodes = inorder(root.Left, inroderNodes)
		inroderNodes = append(inroderNodes, root.Val.(int))
		inroderNodes = inorder(root.Right, inroderNodes)
		return inroderNodes
	}

	inroderNodes := inorder(root, []int{})
	fmt.Println(inroderNodes)

	return inroderNodes[k-1]
}
