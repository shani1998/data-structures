package questions

import "github.com/shani1998/data-structures/tree"

/*
https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/


Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.


Example 1:
        3
       / \
      1   4
       \
        2
Input: root = [3,1,4,null,2], k = 1
Output: 1
*/



func inOrder(root *tree.Node, result *[]tree.Node) {
    if root == nil {
        return
    }
    inOrder(root.Left, result)
    *result = append(*result, *root)
    inOrder(root.Right, result)
}

// time complexity: O(n), space complexity: O(n)
func kthSmallest(root *tree.Node, k int) int {
    inOrderNodes := []tree.Node{}   // Allocate slice
    inOrder(root, &inOrderNodes)   // Pass pointer to slice

    return inOrderNodes[k-1].Val.(int)   // Correct indexing
}

// time complexity: O(h + k), space complexity: O(h), where h is the height of the tree
func kthSmallestOptimized(root *tree.Node, k int) int {
	stack := []*tree.Node{}
	current := root
	count := 0
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		if count == k {
			return current.Val.(int)
		}
		current = current.Right
	}
	return -1 // This line should never be reached if k is valid
}