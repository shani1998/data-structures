package questions

import "github.com/shani1998/data-structures/tree"

/*
https://leetcode.com/problems/cousins-in-binary-tree/description/
Given the root of a binary tree with unique values and the values of two different nodes of the tree x and y,
return true if the nodes corresponding to the values x and y in the tree are cousins, or false otherwise.

Two nodes of a binary tree are cousins if they have the same depth with different parents.

Note that in a binary tree, the root node is at the depth 0, and children of each depth k node are at the depth k + 1.
*/

func getDepthAndParent(root *tree.Node, target int, depth int, parent *tree.Node) (int, *tree.Node) {
    if root == nil {
        return -1, nil
    }

    // Found the node
    if root.Val == target {
        return depth, parent
    }

    // Search left subtree
    d, p := getDepthAndParent(root.Left, target, depth+1, root)
    if d != -1 {
        return d, p
    }

    // Search right subtree
    return getDepthAndParent(root.Right, target, depth+1, root)
}

// time complexity: O(n), space complexity: O(n)
func isCousinsDfs(root *tree.Node, x int, y int) bool {
    dx, px := getDepthAndParent(root, x, 0, nil)
    dy, py := getDepthAndParent(root, y, 0, nil)

    // Same depth, different parents
    return dx == dy && px != py
}

// time complexity: O(n), space complexity: O(n)
 func isCousinsBfs(root *tree.Node, x int, y int) bool {
    if root == nil {
        return false
    }

    queue := []*tree.Node{root}

    for len(queue) > 0 {
        levelSize := len(queue)
        foundX, foundY := false, false
        var parentX, parentY *tree.Node

        for i := 0; i < levelSize; i++ {
            node := queue[i]

            // Check left child
            if node.Left != nil {
                if node.Left.Val == x {
                    foundX = true
                    parentX = node
                }
                if node.Left.Val == y {
                    foundY = true
                    parentY = node
                }
                queue = append(queue, node.Left)
            }

            // Check right child
            if node.Right != nil {
                if node.Right.Val == x {
                    foundX = true
                    parentX = node
                }
                if node.Right.Val == y {
                    foundY = true
                    parentY = node
                }
                queue = append(queue, node.Right)
            }
        }

        // Remove processed level
        queue = queue[levelSize:]

        // If both found in same level → check if parents differ
        if foundX && foundY {
            return parentX != parentY
        }

        // If only one found → cannot be cousins
        if foundX || foundY {
            return false
        }
    }

    return false
}