package questions

import "github.com/shani1998/data-structures/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *tree.Node) []any {
	if root == nil {
		return []any{}
	}

	var result []any
	queue := []*tree.Node{root}

	for len(queue) > 0 {
		levelLen := len(queue)

		for i := 0; i < levelLen; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add left and right children to the queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			// Add the value of the last node in the current level
			if i == levelLen-1 {
				result = append(result, node.Val)
			}
		}
	}

	return result
}
