package tree

type BST struct {
	root *Node
}

func (bst *BST) Insert(Val int) {
	bst.insertHelper(bst.root, Val)
}

func (bst *BST) insertHelper(node *Node, Val int) *Node {
	if bst.root == nil {
		bst.root = &Node{Val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{Val, nil, nil}
	}

	if Val <= node.Val.(int) {
		node.Left = bst.insertHelper(node.Left, Val)
	}

	if Val > node.Val.(int) {
		node.Right = bst.insertHelper(node.Right, Val)
	}

	return node
}
