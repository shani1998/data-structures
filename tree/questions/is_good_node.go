package questions

import "github.com/shani1998/data-structures/tree"

/*
https://leetcode.com/problems/count-good-nodes-in-binary-tree/

Given a binary tree root, a node X in the tree is named "good" if in the path from root to X there are no nodes with a value greater than X.
Return the number of good nodes in the binary tree.

Example 1:
		
*/
// goodNodes returns the number of "good" nodes in the binary tree.
// A node is "good" if its value is >= every value on the path from root to that node.
func goodNodes(root *tree.Node) int {
    if root == nil {
        return 0
    }
    return countGood(root, root.Val.(int))
}

// countGood recursively counts good nodes.
// node      = current node
// maxSoFar  = maximum value seen from the root to this node
// time complexity: O(n) space complexity: O(h) where h is the height of the tree
func countGood(node *tree.Node, maxSoFar int) int {
    if node == nil {
        return 0
    }

    // Check if current node is good
    isGood := 0
    if node.Val.(int) >= maxSoFar {
        isGood = 1
        maxSoFar = node.Val.(int) // update path maximum
    }

    // Count in left and right subtree
    leftCount := countGood(node.Left, maxSoFar)
    rightCount := countGood(node.Right, maxSoFar)

    return isGood + leftCount + rightCount
}