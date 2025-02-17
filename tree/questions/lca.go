package questions

import "github.com/shani1998/data-structures/tree"

func lowestCommonAncestor(root, p, q *tree.Node) *tree.Node {
	pathNodeP := getPath(root, p) // [3,5]
	pathNodeQ := getPath(root, q) // [3,5,2,4]

	minLen := min(len(pathNodeP), len(pathNodeQ))
	var lcaNode *tree.Node
	for i := 0; i < minLen; i++ {
		if pathNodeP[i] == pathNodeQ[i] {
			lcaNode = pathNodeP[i]
		}
	}

	return lcaNode
}

// return path from root to target node
func getPath(root, targetNode *tree.Node) []*tree.Node {
	var dfs func(*tree.Node, *tree.Node, []*tree.Node, *bool)
	var result []*tree.Node

	dfs = func(root, targetNode *tree.Node, currPath []*tree.Node, isPathFound *bool) {
		if root == nil || *isPathFound {
			return
		}
		currPath = append(currPath, root)
		if root == targetNode {
			result = append([]*tree.Node{}, currPath...) // Make a copy of the path
			*isPathFound = true
			return
		}

		dfs(root.Left, targetNode, currPath, isPathFound)
		dfs(root.Right, targetNode, currPath, isPathFound)
	}

	isPathFound := false
	dfs(root, targetNode, []*tree.Node{}, &isPathFound)
	return result
}
