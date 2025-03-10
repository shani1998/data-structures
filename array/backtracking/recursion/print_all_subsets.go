package recursion

// subsets generates all subsets of the given array.
// time complexity: O(n*2^n) , space complexity: O(n*2^n)
func subsets(nums []int) [][]int {
	var res [][]int
	var subset []int
	var printAll func(int)

	// printAll is a recursive function that generates all subsets of the given array.
	printAll = func(i int) {
		if i == len(nums) {
			cpy := make([]int, len(subset))
			copy(cpy, subset)
			res = append(res, cpy)
			return
		}
		subset = append(subset, nums[i]) // Include the current element
		printAll(i + 1)                  // Move to the next element
		subset = subset[:len(subset)-1]  // Exclude the current element
		printAll(i + 1)                  // Move to the next element
	}

	printAll(0)
	return res
}

// time to generate all subsets would be O(2^n) but copy operation will take O(n) time so total time complexity would be O(n*2^n)
func generateAllSubsets(start int, nums, subset []int, subsets [][]int) [][]int {
	if start == len(nums) {
		tmpArr := make([]int, len(subset))
		copy(tmpArr, subset)
		subsets = append(subsets, tmpArr)
		return subsets
	}

	// include the current element
	subset = append(subset, nums[start])
	subsets = generateAllSubsets(start+1, nums, subset, subsets)

	// exclude the current element
	subset = subset[:len(subset)-1]
	subsets = generateAllSubsets(start+1, nums, subset, subsets)

	return subsets
}

/*
                []
               /   \
            [1]     []
           /   \     /   \
      [1,2]    [1]  [2]    []
      /   \    / \    / \    / \
 [123] [12]  [13][1] [23][2] [3][]

*/
