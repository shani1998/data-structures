package tree

type BST struct {
	root *Node
}

func (bst *BST) Insert(val int) {
	bst.insertHelper(bst.root, val)
}

func (bst *BST) insertHelper(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}

	if val <= node.val.(int) {
		node.left = bst.insertHelper(node.left, val)
	}

	if val > node.val.(int) {
		node.right = bst.insertHelper(node.right, val)
	}

	return node
}
