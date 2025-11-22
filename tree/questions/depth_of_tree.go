package questions

import (
	"github.com/shani1998/data-structures/tree"
)

/*
Given the root of a binary tree, return its maximum depth.
A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
Example 1:
Input: root = [3,9,20,null,null,15,7]
Output: 3
*/

func maxDepth(root *tree.Node) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepthPath(root *tree.Node) []int {
	return maxDepthPathUtil(root, []int{})
}

func maxDepthPathUtil(root *tree.Node, path []int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		path = append(path, root.Val.(int))
		return path
	}

	path = append(path, root.Val.(int))
	leftPath := maxDepthPathUtil(root.Left, path)
	rightPath := maxDepthPathUtil(root.Right, path)
	if len(leftPath) > len(rightPath) {
		return leftPath
	}
	return rightPath
}

func maxDepthPath1(root *tree.Node) []int {
	if root == nil {
		return []int{}
	}

	// Get the path for the left and right subtrees
	leftPath := maxDepthPath(root.Left)
	rightPath := maxDepthPath(root.Right)

	// Choose the longer path and append the current node's value
	if len(leftPath) > len(rightPath) {
		return append(leftPath, root.Val.(int))
	}
	return append(rightPath, root.Val.(int))
}

func brokenCalc(startValue int, target int) int {
	var isTargetFound bool
	return brokenCalcUtil(startValue, startValue, target, 0, &isTargetFound)
}

func brokenCalcUtil(currVal, startVal, target int, countOpn int, isTargetFound *bool) int { // 2, 2, 3, 0
	if currVal <= 0 || currVal > target || *isTargetFound {
		return countOpn
	}
	if currVal == target {
		*isTargetFound = true
		return countOpn
	}

	return min(brokenCalcUtil(currVal*2, startVal, target, countOpn+1, isTargetFound),
		brokenCalcUtil(currVal-1, startVal, target, countOpn+1, isTargetFound))
}

/*
Given the root of a binary tree, return the length of the diameter of the tree.
The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
This path may or may not pass through the root.
The length of a path between two nodes is represented by the number of edges between them.

Example:
       1
      / \
     2   3
    / \
   4   5
Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
*/

func diameterOfBinaryTree(root *tree.Node) int {
	maxDia := 0
	depth(root, &maxDia)
	return maxDia
}

func depth(node *tree.Node, maxDia *int) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left, maxDia)
	right := depth(node.Right, maxDia)

	if left+right > *maxDia {
		*maxDia = left + right
	}

	return 1 + max(left, right)
}
