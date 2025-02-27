package main

import "math"

// Write an efficient program to find the sum of contiguous subarray
// within a one-dimensional array of numbers that has the largest sum.
/*
Use the input vector nums to store the candidate subarrays sum (i.e. the greatest contiguous sum so far).
Ignore cumulative negatives, as they don't contribute positively to the sum.
Ex:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
*/

// Time complexity: O(n*n), Space complexity: O(1)
func maxSubArray1(nums []int) int {
	maxSum := math.MinInt64

	for i := 0; i < len(nums); i++ {
		currSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			currSum += nums[j]
			maxSum = max(maxSum, currSum)
		}
	}

	return maxSum
}

// Time complexity: O(n), Space complexity: O(1)
func maxSubArray(nums []int) int {
	maxSum, currSum := nums[0], 0 // or maxSumSoFar, currSum := math.MinInt, 0

	for _, num := range nums {
		currSum += num
		maxSum = max(maxSum, currSum)
		if currSum <= 0 {
			currSum = 0
		}
	}

	return maxSum
}

// return left and right index of the sub array [-2,1,-3,4,-1,2,1,-5,4] -> [4,-1,2,1]
func maxSubArrayIndices(nums []int) (maxLeft int, maxRight int) {
	maxSum, currSum := nums[0], 0
	currLeft := 0

	for i, num := range nums {
		currSum += num
		if currSum > maxSum {
			maxSum = currSum
			maxLeft, maxRight = currLeft, i
		}
		if currSum <= 0 {
			currSum = 0
			currLeft = i + 1
		}

	}

	return maxLeft, maxRight
}

// maximum-absolute-sum-of-any-subarray Ex: nums = [2,-5,1,-4,3,-2] -> 8
// time complexity: O(2n), space complexity: O(1)
func maxAbsoluteSum1(nums []int) int {
	maxSum, minSum, currSum := 0, 0, 0

	// maximum positive sum
	for _, num := range nums {
		currSum += num
		maxSum = max(maxSum, currSum)
		if currSum < 0 {
			currSum = 0
		}
	}

	// maximum negative sum
	currSum = 0
	for _, num := range nums {
		currSum += num
		minSum = min(minSum, currSum)
		if currSum > 0 {
			currSum = 0
		}
	}

	return max(maxSum, -minSum)
}

// Time complexity: O(n), Space complexity: O(1)
func maxAbsoluteSum2(nums []int) int {
	maxSum, minSum := 0, 0
	currMax, currMin := 0, 0

	for _, num := range nums {
		currMax += num
		currMin += num

		maxSum = max(maxSum, currMax)
		minSum = min(minSum, currMin)

		if currMax < 0 {
			currMax = 0
		}
		if currMin > 0 {
			currMin = 0
		}
	}

	return max(maxSum, -minSum)
}
