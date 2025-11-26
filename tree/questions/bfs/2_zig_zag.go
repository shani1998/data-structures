package bfs

import "github.com/shani1998/data-structures/tree"

/*
https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/description/

Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
(i.e., from left to right, then right to left for the next level and alternate between).

Example 1:
        3
      /   \
     9     20
          /  \
         15   7
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[20,9],[15,7]]
*/

// time complexity: O(n), space complexity: O(n)
func zigzagLevelOrder(root *tree.Node) [][]int {
    if root == nil {
        return [][]int{}
    }

    queue := []*tree.Node{root}
    front := 0
    leftToRight := true
    result := [][]int{}

    for front < len(queue) {
        levelSize := len(queue) - front
        level := make([]int, levelSize)

        for i := 0; i < levelSize; i++ {
            node := queue[front]
            front++

            // Fill left->right or right->left
            if leftToRight {
                level[i] = node.Val.(int)
            } else {
                level[levelSize-1-i] = node.Val.(int)
            }

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, level)
        leftToRight = !leftToRight
    }

    return result
}
