package questions

func twoSumUnsortedArr(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if atIndex, ok := numsMap[target-num]; ok {
			return []int{atIndex, i}
		}
		numsMap[num] = i
	}
	return []int{}
}

func twoSumSortedArr(numbers []int, target int) []int {
	totalNums := len(numbers)
	left, right := 0, totalNums-1

	for left < right {
		currSum := numbers[left] + numbers[right]
		if currSum == target {
			return []int{left + 1, right + 1} // return answer in 1th index
		}

		// if currSum is less than target then increase left by 1
		if currSum < target {
			left++
			continue
		}
		right--
	}

	return []int{}
}

/*
You are given a 0-indexed array nums consisting of positive integers. You can choose two indices i and j, such that i != j, an
d the sum of digits of the number nums[i] is equal to that of nums[j].
Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and j that satisfy the conditions.

Input: nums = [18,43,36,13,7]
Output: 54
Explanation: The pairs (i, j) that satisfy the conditions are:
- (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
- (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
So the maximum sum that we can obtain is 54.
*/
func maximumSum(nums []int) int {
	sumDigitMap := make(map[int]int) // Store the largest number for each digit sum
	maxSum := -1

	for _, num := range nums {
		digitSum := sumDigit(num)

		if largestNum, exists := sumDigitMap[digitSum]; exists {
			// If a number with the same digit sum exists, calculate sum
			maxSum = max(maxSum, num+largestNum)

			// Update the largest number if the current num is greater
			if num > largestNum {
				sumDigitMap[digitSum] = num
			}
		} else {
			// If no number with this digit sum exists, store the current num
			sumDigitMap[digitSum] = num
		}
	}

	return maxSum
}

func sumDigit(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
