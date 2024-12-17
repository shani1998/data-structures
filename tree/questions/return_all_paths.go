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
