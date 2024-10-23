package questions

import "github.com/shani1998/data-structures/tree"

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
