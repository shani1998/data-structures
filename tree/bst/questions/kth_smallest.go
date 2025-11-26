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

func kthSmallest(root *tree.Node, k int) int {
    inOrderNodes := []tree.Node{}   // Allocate slice
    inOrder(root, &inOrderNodes)   // Pass pointer to slice

    return inOrderNodes[k-1].Val.(int)   // Correct indexing
}