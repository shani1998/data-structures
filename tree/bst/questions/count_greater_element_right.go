package questions

/*
Given an array A with n values, let X of A be an array that holds in index i the number of elements which are bigger than A[i] and are to its right side in the original array A.

For example, if A was: [10,12,8,17,3,24,19], then X(A) is: [4,3,3,2,2,0,0]

*/

// CountGreaterElementRight returns the count of elements greater than the current element to the right
// time comlexity: O(n^2), space: O(1)
func CountGreaterElementRight(nums []int) []int {
	totalNum := len(nums)
	result := make([]int, totalNum)

	// Iterate twice to handle circular array
	for i := 0; i < totalNum; i++ {
		for j := i + 1; j < totalNum; j++ {
			if nums[j] > nums[i] {
				result[i]++
			}
		}
	}

	return result
}

//
//func CountGreaterElementRight1(nums []int) []int {
//	lenNum := len(nums)
//
//	// traverse from right and keep inserting elements to BST, for each element
//	// use the binary search tree's find() operation to find how many elements are greater than e in the current tree.
//	for i := lenNum - 1; i >= 0; i-- {
//
//	}
//}

// TreeNode represents a node in the BST.
type TreeNode struct {
	val       int
	rightSize int // Number of elements in the right subtree
	count     int // Number of times this value appears
	left      *TreeNode
	right     *TreeNode
}

// Insert a new value into the BST and count greater elements to the right.
func insert(root *TreeNode, val int) (newRoot *TreeNode, count int) {
	if root == nil {
		return &TreeNode{val: val, count: 1}, 0
	}

	if val < root.val {
		// Go left; these elements are NOT greater
		newLeft, cnt := insert(root.left, val)
		root.left = newLeft
		return root, cnt
	} else {
		// All elements in left + current node count as greater elements
		count = root.rightSize + root.count
		if val == root.val {
			root.count++
		} else {
			newRight, cnt := insert(root.right, val)
			root.right = newRight
			count += cnt
		}
		root.rightSize++ // Increase right subtree size
		return root, count
	}
}

// Function to compute X(A) using BST
func countGreaterRight(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	var root *TreeNode

	// Traverse the array from right to left
	for i := n - 1; i >= 0; i-- {
		var count int
		root, count = insert(root, arr[i])
		result[i] = count
	}

	return result
}
