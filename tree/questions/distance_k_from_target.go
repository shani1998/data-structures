package questions

import "github.com/shani1998/data-structures/tree"

/*
https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

We are given a binary tree (with root node root), a target node, and an integer value k.

Return a list of the values of all nodes that have a distance k from the target node.  The answer can be returned in any order.

Example 1:
Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
Output: [7,4,1]
Explanation: The nodes that are a distance 2 from the target node (with value 5)
have values 7, 4, and 1.
*/

// time complexity: O(n) space complexity: O(n)
func distanceK(root, target *tree.Node, k int) []int {
    // Step 1: Build graph adjacency map
    graph := make(map[*tree.Node][]*tree.Node)
    buildGraph(root, nil, graph)

    // Step 2: BFS from target
    queue := []*tree.Node{target}
    visited := map[*tree.Node]bool{target: true}
    dist := 0

    for len(queue) > 0 {
        size := len(queue)
        if dist == k {
            // Collect values of all nodes currently in the queue
            result := make([]int, size)
            for i := 0; i < size; i++ {
                result[i] = queue[i].Val.(int)
            }
            return result
        }

        // Expand BFS
        for i := 0; i < size; i++ {
            node := queue[i]
            for _, nei := range graph[node] {
                if !visited[nei] {
                    visited[nei] = true
                    queue = append(queue, nei)
                }
            }
        }

        queue = queue[size:] // pop level
        dist++
    }

    return []int{}
}

// Build undirected graph from the tree
// time complexity: O(n) space complexity: O(n)
func buildGraph(node, parent *tree.Node, graph map[*tree.Node][]*tree.Node) {
    if node == nil {
        return
    }

    if parent != nil {
        graph[node] = append(graph[node], parent)
        graph[parent] = append(graph[parent], node)
    }

    buildGraph(node.Left, node, graph)
    buildGraph(node.Right, node, graph)
}