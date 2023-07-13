package hashing

import "sort"

// LongestConsecutive returns the length of the longest consecutive elements sequence, https://leetcode.com/problems/longest-consecutive-sequence/
func longestConsecutiveSol1(nums []int) int { // time: O(nlogn), space: O(1)
	sort.Ints(nums) // O(nlogn)

	totalNum := len(nums)
	if totalNum <= 1 {
		return totalNum
	}

	maxConsecutiveSoFar, output := 1, 0
	for i := 0; i < totalNum-1; i++ {
		if nums[i] == nums[i+1] {
			if maxConsecutiveSoFar > output {
				output = maxConsecutiveSoFar
			}
			continue

		}
		if nums[i+1] == nums[i]+1 {
			maxConsecutiveSoFar++
			if maxConsecutiveSoFar > output {
				output = maxConsecutiveSoFar
			}
			continue
		}
		maxConsecutiveSoFar = 1
	}

	return output
}

func longestConsecutiveSol2(nums []int) int { // time: O(n^2), space: O(n)
	lookup := make(map[int]struct{})
	for _, num := range nums {
		lookup[num] = struct{}{}
	}

	totalNums, outputLen := len(nums), 0
	for i := 0; i < totalNums; i++ {
		if _, ok := lookup[nums[i]-1]; ok {
			continue // as this won't be a start of consecutive window
		}

		maxLenSoFar := 1 // assume it's start of the sequence

		//traverse until nums[i]+1 found in hash
		currNum, isFound := nums[i], true
		for isFound {
			if _, isFound = lookup[currNum+1]; isFound {
				currNum++
				maxLenSoFar++
			}
			if maxLenSoFar > outputLen {
				outputLen = maxLenSoFar
			}

		}
	}

	return outputLen
}
