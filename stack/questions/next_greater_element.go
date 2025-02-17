package questions

// NextGreaterElement returns the next greater element of the input array.
// The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.
// If there is no such element, set 0.
/*
Ex:
Input: [4,1,2,1,3]
Output: [0,2,3,3,0]
*/

// NextGreaterElement1 Brute force approach: time complexity O(n^2), space complexity: O(n)
func NextGreaterElement1(nums []int) []int {
	totalNum := len(nums)
	result := make([]int, totalNum)

	for i := 0; i < totalNum; i++ {
		result[i] = 0
		for j := i + 1; j < totalNum; j++ {
			if nums[j] > nums[i] {
				result[i] = nums[j]
				break
			}
		}
	}

	return result
}

// NextGreaterElement2 using stack: time complexity O(n), space complexity O(n)
func NextGreaterElement2(nums []int) []int {
	totalNum := len(nums)
	stack := make([]int, 0)
	result := make([]int, totalNum)

	// Iterate twice to handle circular array
	for i := totalNum - 1; i >= 0; i-- {

		// Maintain decreasing stack, pop elements smaller than current number
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1] // pop
		}

		// If stack is not empty, the top of the stack is the next greater element
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		} else {
			result[i] = 0
		}

		// Push the current index onto the stack
		stack = append(stack, nums[i])
	}

	return result
}
