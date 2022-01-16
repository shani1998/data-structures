package main

// Write an efficient program to find the sum of contiguous subarray
// within a one-dimensional array of numbers that has the largest sum.
/*
Use the input vector nums to store the candidate subarrays sum (i.e. the greatest contiguous sum so far).
Ignore cumulative negatives, as they don't contribute positively to the sum.
*/
func maxSubArray(nums []int) int {
	maxSumSoFar, currSum := -1000000, 0
	for _, v := range nums {
		currSum += v
		if currSum > maxSumSoFar {
			maxSumSoFar = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSumSoFar
}
