package slidingwindow

import "math"

/*
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
A subarray is a contiguous non-empty sequence of elements within an array.
Ex1:
Input: nums = [1,1,1], k = 2
Output: 2

Ex2:
Input: nums = [1,2,3], k = 3
Output: 2
*/

// subArraySumEqualToK1 is a brute force solution, time complexity: O(n^2)
func subArraySumEqualToK1(nums []int, k int) int {
	count, totalNums := 0, len(nums)
	for i := 0; i < totalNums; i++ {
		sum := 0
		for j := i; j < totalNums; j++ {
			sum += nums[i]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// subArraySumEqualToK2 is a more optimized solution, time complexity: O(n)
func subArraySumEqualToK2(nums []int, k int) int {
	prefixSum := make(map[int]int) // Map to store prefix sums and their frequencies
	sum, count := 0, 0

	for _, num := range nums {
		sum += num
		if sum == k {
			count++
		}

		// Check if there's a subarray ending at the current index with sum equal to k
		if freq, ok := prefixSum[sum-k]; ok {
			count += freq
		}

		// Update the map with the current prefix sum
		prefixSum[sum]++
	}

	return count
}

// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.
// This works when all the elements are positive using sliding window.
func subarraySum(nums []int, k int) int {
	left, right, windowSum := 0, 0, 0
	totalCount, n := 0, len(nums)

	for right < n {
		// expand window by adding nums[right]
		windowSum += nums[right]
		right++

		// shrink window while sum > k
		for windowSum > k && left < right {
			windowSum -= nums[left]
			left++
		}

		// check if we found a subarray
		if windowSum == k {
			totalCount++
		}
	}

	return totalCount
}

// longestSubArraySumEqualTok, returns the length of longest subarray which sum equals to k
func longestSubArraySumEqualTok(arr []int, k int) int {
	prefixSum := make(map[int]int)
	sum, maxLen := 0, 0

	for i, num := range arr {
		sum += num

		if indexAt, ok := prefixSum[sum-k]; ok {
			maxLen = max(maxLen, i-indexAt+1)
		}

		if _, ok := prefixSum[sum]; !ok {
			prefixSum[sum] = i
		}

	}

	return maxLen
}

/*
MaxSubarrayLenK
"Find the maximum sum of sub-array of size k with the time complexity of O(N).
Array = [1, 2, 6, 2, 4, 1], k = 3"
*/
func MaxSubarrayLenK(arr []int, k int) int { // O(n*k)
	n, maxSum := len(arr), 0
	for i := 0; i < n-k+1; i++ { // n-k iterations ~ n
		step, sum := i+k, 0
		for j := i; j < step; j++ { // k iterations
			sum += arr[j]
		}
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

/*
MaxSubarrayLenK1
instead of re-calculating the window everytime ... we can reuse the sum
one number is getting added and one is getting subtracted.
*/
func MaxSubarrayLenK1(arr []int, k int) int {
	windowSum, maxSum := 0, 0
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	left, right, n := 0, k, len(arr)
	for right < n {
		windowSum += arr[right] - arr[left]
		maxSum = max(maxSum, windowSum)
		left++
		right++
	}

	return maxSum
}

/*
SmallestSubarraySumGtEqK:
given an array of positive integers nums and a target sum target,
find the length of the smallest contiguous subarray whose sum is greater than or equal to target.
If no such subarray exists, return 0.
Input: nums = [2, 3, 1, 2, 4, 3], target = 7
Output: 2
Explanation: The subarray [4, 3] has the smallest length = 2
*/

func SmallestSubarraySumGtEqK(arr []int, k int) int {
	left, right, currSum := 0, 0, arr[0]
	minLen, n := math.MaxInt64, len(arr)

	for right < n {
		if currSum >= k {
			minLen = min(minLen, right-left+1)
			currSum -= arr[left]
			left++ // shrink the window

		}
		if currSum < k { // expand the window, as we have positive no. in arr
			right++
			if right < n {
				currSum += arr[right]
			}

		}
	}
	return minLen
}
