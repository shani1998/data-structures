package searching

import "math"

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1 // Handle empty array case
	}

	low, high := 0, n-1
	minNum := math.MaxInt

	for low <= high {
		mid := (low + high) / 2

		// Always update the minimum value found
		minNum = min(minNum, nums[mid])

		// If the left half is sorted
		if nums[low] <= nums[mid] {
			// The minimum could be nums[low], so update minNum
			minNum = min(minNum, nums[low])
			// Discard the left half
			low = mid + 1
		} else { // Otherwise, right half is sorted
			// The minimum could be nums[mid + 1], so update minNum
			minNum = min(minNum, nums[mid+1])
			// Discard the right half
			high = mid - 1
		}
	}

	return minNum
}
