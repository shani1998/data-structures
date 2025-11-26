package bfs

import "github.com/shani1998/data-structures/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *tree.Node) []int {
    if root == nil {
        return []int{}
    }

    queue := []*tree.Node{root}
    result := []int{}

    for len(queue) > 0 {
        size := len(queue)
        var lastVal int

        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]

            // track the last node of this level
            lastVal = node.Val.(int)

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }

        result = append(result, lastVal)
    }

    return result
}
