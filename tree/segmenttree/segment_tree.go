package main

import "fmt"

type SegmentTree struct {
	node []int
	n    int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	return &SegmentTree{
		node: make([]int, 4*n),
		n:    n,
	}
}

// Build builds the segment tree
func (tree *SegmentTree) build(arr []int, index, left, right int) {
	if left == right {
		tree.node[index] = arr[left]
		return
	}

	middle := (left + right) / 2

	tree.build(arr, index*2+1, left, middle)    // left child would be at index*2 + 1
	tree.build(arr, index*2+2, middle+1, right) // right child would be at index*2 + 2
	tree.node[index] = tree.node[index*2+1] + tree.node[index*2+2]
}

// Query returns the sum of elements in the range [queryLeft, queryRight]
func (tree *SegmentTree) query(index, left, right, queryLeft, queryRight int) int {
	// If the range of the query is outside the range of the current node
	if queryLeft > right || queryRight < left {
		return 0
	}
	// If the range of the query is inside the range of the current node
	if queryLeft <= left && right <= queryRight {
		return tree.node[index]
	}

	middle := (left + right) / 2
	leftSum := tree.query(index*2+1, left, middle, queryLeft, queryRight)
	rightSum := tree.query(index*2+2, middle+1, right, queryLeft, queryRight)
	return leftSum + rightSum
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	tree := NewSegmentTree(arr)
	tree.build(arr, 0, 0, tree.n-1)
	fmt.Println(tree.node)

	fmt.Println(tree.query(0, 0, tree.n-1, 1, 3)) // 15
}
