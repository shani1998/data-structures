package questions

import "fmt"

// Node of the perfect binary tree
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// BuildPerfectTree builds a perfect binary tree of height h
// returns root + slice of leaf nodes
func BuildPerfectTree(h int) (*Node, []*Node) {
	if h == 0 {
		return nil, nil
	}

	root := &Node{}
	leaves := []*Node{}

	var build func(level int) *Node
	build = func(level int) *Node {
		n := &Node{}

		if level == h {
			leaves = append(leaves, n)
			return n
		}

		n.Left = build(level + 1)
		n.Right = build(level + 1)

		n.Left.Parent = n
		n.Right.Parent = n

		return n
	}

	root = build(1)
	return root, leaves
}

// Fix up values from a leaf to root
func fixUp(n *Node) {
	curr := n.Parent

	for curr != nil {
		leftVal := curr.Left.Val
		rightVal := curr.Right.Val
		curr.Val = leftVal & rightVal
		curr = curr.Parent
	}
}

// Set a leaf value to 1
func Set(leaf *Node) {
	leaf.Val = 1
	fixUp(leaf)
}

// Clear a leaf value to 0
func Clear(leaf *Node) {
	leaf.Val = 0
	fixUp(leaf)
}

// Print tree (for debugging)
func PrintTree(root *Node, level int) {
	if root == nil {
		return
	}
	PrintTree(root.Right, level+1)
	fmt.Printf("%s%d\n", indent(level), root.Val)
	PrintTree(root.Left, level+1)
}

func indent(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "  "
	}
	return s
}

func main() {
	// Build perfect tree of height 3
	root, leaves := BuildPerfectTree(3)

	// Initial (all 0)
	fmt.Println("Initial tree:")
	PrintTree(root, 0)

	// Set leaf #1 and leaf #3
	fmt.Println("\nSet leaf 1 and leaf 3 to 1")
	Set(leaves[1])
	Set(leaves[3])
	PrintTree(root, 0)

	// Set all leaves to 1 → entire tree becomes 1
	fmt.Println("\nSet all leaves to 1")
	for _, l := range leaves {
		Set(l)
	}
	PrintTree(root, 0)

	// Clear one leaf
	fmt.Println("\nClear leaf 2")
	Clear(leaves[2])
	PrintTree(root, 0)
}

// Below is the cleanest and most efficient solution using an array representation of a perfect binary tree (just like a heap).
// This avoids pointers, avoids building a tree, and makes set() / clear() O(log N).
// leaf_start = size_of_tree // 2 as it is a complete bin tree

type BitTree struct {
	tree      []int
	leafStart int
	leafCount int
}

// NewBitTree Build perfect binary tree with leafCount leaves
func NewBitTree(leafCount int) *BitTree {
	// total nodes = 2*leafCount - 1
	size := 2*leafCount - 1
	leafStart := leafCount - 1

	return &BitTree{
		tree:      make([]int, size),
		leafStart: leafStart,
		leafCount: leafCount,
	}
}

// Set leaf i to 1, then fix parents
func (b *BitTree) Set(i int) {
	pos := b.leafStart + i
	b.tree[pos] = 1
	b.fixUp(pos)
}

// Clear leaf i to 0, then fix parents
func (b *BitTree) Clear(i int) {
	pos := b.leafStart + i
	b.tree[pos] = 0
	b.fixUp(pos)
}

// Fix parents using rule: parent = left & right
func (b *BitTree) fixUp(pos int) {
	for pos > 0 {
		parent := (pos - 1) / 2
		left := 2*parent + 1
		right := left + 1

		// AND rule
		b.tree[parent] = b.tree[left] & b.tree[right]

		pos = parent
	}
}

func (b *BitTree) Root() int {
	return b.tree[0]
}

func main1() {
	bt := NewBitTree(4) // 4 leaves → 7 nodes

	bt.Set(0)
	bt.Set(1)
	bt.Set(2)
	bt.Set(3)

	fmt.Println("Root after all set:", bt.Root()) // 1

	bt.Clear(2)

	fmt.Println("Root after clearing leaf2:", bt.Root()) // 0
}
