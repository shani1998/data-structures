package questions

import (
	"math"

	"github.com/shani1998/data-structures/tree"
)

func getInorderTraversal(root *tree.Node, result []int) []int {
	if root == nil {
		return result
	}

	result = getInorderTraversal(root.Left, result)
	result = append(result, root.Val.(int))
	result = getInorderTraversal(root.Right, result)

	return result
}

func isValidBST(root *tree.Node) bool {
	// inorder traversal of BST should be sorted array
	traversal := getInorderTraversal(root, []int{})

	// check is array is sorted or not
	n := len(traversal)
	for i := 1; i < n; i++ {
		if traversal[i-1] >= traversal[i] {
			return false
		}
	}
	return true
}

func isValidBST1(root *tree.Node) bool {
	return validate(root, math.MinInt, math.MaxInt)
}

func validate(root *tree.Node, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val.(int) <= min || root.Val.(int) >= max {
		return false
	}

	return validate(root.Left, min, root.Val.(int)) && validate(root.Right, root.Val.(int), max)
}
