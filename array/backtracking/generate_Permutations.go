package backtracking

/*
GeneratePermutations generates all permutations of the input array.
Ex:
Input: [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func GeneratePermutations(nums []int) [][]int {
	return permute(nums, 0, [][]int{})
}

func permute(nums []int, start int, result [][]int) [][]int {
	if start == len(nums)-1 {
		// Make a copy of nums to avoid modifying the appended slice later
		copiedNums := make([]int, len(nums))
		copy(copiedNums, nums)
		result = append(result, copiedNums)
		return result
	}

	for i := start; i < len(nums); i++ {
		nums[i], nums[start] = nums[start], nums[i]
		result = permute(nums, start+1, result)
		nums[i], nums[start] = nums[start], nums[i]
	}

	return result
}
