package questions

import "sort"

/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
*/

// Time complexity: O(nlogn), Space complexity: O(n)
func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n <= 1 {
		return intervals
	}

	// Sort intervals based on the start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}
	prevInt := intervals[0]

	for i := 1; i < n; i++ {
		currInt := intervals[i]

		// Check if intervals overlap
		if currInt[0] <= prevInt[1] {
			// Merge the intervals
			prevInt[1] = max(prevInt[1], currInt[1])
		} else {
			// Add the non-overlapping interval to the result
			result = append(result, prevInt)
			prevInt = currInt
		}
	}

	// Add the last interval
	result = append(result, prevInt)

	return result
}
