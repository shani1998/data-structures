package tree

func getInorderTraversal(root *Node, result []int) []int {
	if root == nil {
		return result
	}

	result = getInorderTraversal(root.left, result)
	result = append(result, root.val.(int))
	result = getInorderTraversal(root.right, result)

	return result
}

func isValidBST(root *Node) bool {
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
