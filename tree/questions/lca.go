package questions

import "github.com/shani1998/data-structures/tree"

func lowestCommonAncestor(root, p, q *tree.Node) *tree.Node {
	pathNodeP := getPath(root, p) // [3,5]
	pathNodeQ := getPath(root, q) // [3,5,2,4]

	// Find the last common node
	var lcaNode *tree.Node
	minLength := min(len(pathNodeP), len(pathNodeQ))

	for i := 0; i < minLength; i++ {
		if pathNodeP[i] == pathNodeQ[i] {
			lcaNode = pathNodeP[i]
		} else {
			break
		}
	}

	return lcaNode
}

// return path from root to target node
func getPath(root, target *tree.Node) []*tree.Node {
	var path []*tree.Node
	dfs(root, target, &path)
	return path
}

func dfs(root, target *tree.Node, path *[]*tree.Node) bool {
	if root == nil {
		return false
	}

	// Add this node to path
	*path = append(*path, root)

	// Target found
	if root == target {
		return true
	}

	// Search left or right
	if dfs(root.Left, target, path) || dfs(root.Right, target, path) {
		return true
	}

    // Backtrack — remove last entry as target not found in this path
	*path = (*path)[:len(*path)-1]

	return false
}

func lowestCommonAncestor2(root, p, q *tree.Node) *tree.Node {
	// Base: empty tree or we've found p or q
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both sides are non-nil, p and q are in different subtrees -> root is LCA
	if left != nil && right != nil {
		return root
	}

	// Otherwise return the non-nil side (or nil if both nil)
	if left != nil {
		return left
	}
	return right
}
