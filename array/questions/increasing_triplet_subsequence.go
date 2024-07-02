package questions

import "math"

func increasingTriplet(nums []int) bool {
	lenNum := len(nums)
	if lenNum < 3 {
		return false
	}
	firstMinSoFar, secondMinSoFar := nums[0], math.MaxInt64
	for i := 1; i < lenNum; i++ {
		if firstMinSoFar < secondMinSoFar && secondMinSoFar < nums[i] {
			return true
		}
		if nums[i] > firstMinSoFar && nums[i] < secondMinSoFar {
			secondMinSoFar = nums[i]
		}
		if nums[i] < firstMinSoFar {
			firstMinSoFar = nums[i]
		}
	}
	return false
}
