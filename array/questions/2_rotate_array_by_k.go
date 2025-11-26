package questions

/*
https://leetcode.com/problems/rotate-array/description/

Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
*/

// time complexity: O(n), space complexity: O(n)
func rotate(nums []int, k int) {
	n := len(nums)
	tempNums := make([]int, n)
	copy(tempNums, nums)

	for i, num := range tempNums {
		newIdx := (i + k) % n
		nums[newIdx] = num
	}
}

// time complexity: O(n), space complexity: O(1)
// use reverse approach to rotate array in place
func rotateInPlace(nums []int, k int) { // [1,2,3,4,5,6,7], k = 3
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1) // reverse the whole array [7,6,5,4,3,2,1]
	reverse(nums, 0, k-1) // reverse first k elements [5,6,7,4,3,2,1]
	reverse(nums, k, n-1) // reverse remaining elements [5,6,7,1,2,3,4]
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}