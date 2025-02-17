package questions

import "github.com/shani1998/data-structures/tree"

/*
Given an integer array nums where the elements are sorted in ascending order, convert it to a
height-balanced binary search tree.
Input: nums = [-10,-3,0,5,9]
Output: [0,-3,9,-10,null,5]
Explanation: [0,-10,5,null,-3,null,9] is also accepted
*/

func sortedArrayToBST(nums []int) *tree.Node {
	return sortedArrayToBSTUtil(nums, 0, len(nums)-1)
}

func sortedArrayToBSTUtil(nums []int, start, end int) *tree.Node {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &tree.Node{Val: nums[mid]}

	root.Left = sortedArrayToBSTUtil(nums, start, mid-1)
	root.Right = sortedArrayToBSTUtil(nums, mid+1, end)

	return root
}
