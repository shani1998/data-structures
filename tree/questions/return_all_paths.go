package questions

import (
	"fmt"

	"github.com/shani1998/data-structures/tree"
)

/*
Given the root of a binary tree, return all root-to-leaf paths in any order.
A leaf is a node with no children.

Example 1:
         1
	   /   \
	 2      3
	  \
	   5
Input: root = [1,2,3,null,5]
Output: ["1->2->5","1->3"]
*/

func binaryTreePaths(root *tree.Node) []string {
	return binaryTreePathsUtils(root, []string{}, "")
}

func binaryTreePathsUtils(root *tree.Node, result []string, path string) []string {
	if root == nil {
		return result
	}
	if root.Left == nil && root.Right == nil {
		path += fmt.Sprintf("%d", root.Val)
		result = append(result, path)
		return result
	}

	path += fmt.Sprintf("%d", root.Val) + "->"
	result = binaryTreePathsUtils(root.Left, result, path)
	result = binaryTreePathsUtils(root.Right, result, path)

	return result
}

// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum.
// Time complexity O(n), space complexity O(n*m)
func pathSum(root *tree.Node, targetSum int) [][]int {
	var dfs func(*tree.Node, []int, int, int) [][]int
	var result [][]int

	dfs = func(root *tree.Node, currPath []int, currPathSum, targetSum int) [][]int {
		if root == nil {
			return result
		}

		currPath = append(currPath, root.Val.(int))
		currPathSum += root.Val.(int)

		if root.Left == nil && root.Right == nil && currPathSum == targetSum {
			tmpArr := make([]int, len(currPath))
			copy(tmpArr, currPath)
			result = append(result, tmpArr)
			return result
		}

		result = dfs(root.Left, currPath, currPathSum, targetSum)
		result = dfs(root.Right, currPath, currPathSum, targetSum)
		return result
	}

	dfs(root, []int{}, 0, targetSum)

	return result
}
