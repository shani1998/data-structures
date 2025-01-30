package slidingwindow

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
