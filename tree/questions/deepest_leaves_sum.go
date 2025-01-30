package questions

import "github.com/shani1998/data-structures/tree"

// https://leetcode.com/problems/deepest-leaves-sum/description
/*
Given the root of a binary tree, return the sum of values of its deepest leaves.
Example 1:
Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15
*/

// return sum of the last level using BFS, Time: O(n), Space: O(n)
func deepestLeavesSum(root *tree.Node) int {
	if root == nil {
		return 0
	}

	queue := []*tree.Node{root}
	var sum int

	// Perform level order traversal
	for len(queue) > 0 {
		qSize := len(queue)
		sum = 0 // Reset sum for each level

		// Traverse all nodes in the current level
		for i := 0; i < qSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add node's value to the sum of the current level
			sum += node.Val.(int)

			// Enqueue left and right children if they exist
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	// `sum` will contain the sum of the deepest leaves after the loop
	return sum
}

// deepestLeavesSum1 return the sum of all nodes at max depth using DFS, Time : O(n), Space : O(h)
func deepestLeavesSum1(root *tree.Node) int {
	var maxDepth int
	var sumAtMaxDepth int

	// Helper function to perform DFS
	var dfs func(node *tree.Node, depth int)
	dfs = func(node *tree.Node, depth int) {
		if node == nil {
			return
		}

		// If we reach a deeper level, reset the sum and update the max depth
		if depth > maxDepth {
			maxDepth = depth
			sumAtMaxDepth = node.Val.(int)
		} else if depth == maxDepth {
			sumAtMaxDepth += node.Val.(int)
		}

		// Continue DFS for left and right children
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	// Start DFS from the root at depth 0
	dfs(root, 0)

	return sumAtMaxDepth
}
