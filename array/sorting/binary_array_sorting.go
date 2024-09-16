package sorting

/*
n = 5
arr = 1 0 1 1 0
*/

// use 2 pointer approach to sort the binary array
func binSort(arr []int, n int) []int {
	i, j := -1, 0

	for j < n {
		if arr[j] == 0 {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
		j++
	}
	return arr
}

// use counting sort to sort the binary array
func binSort1(arr []int, n int) []int {
	// TODO implement counting sort
	return arr
}

/*
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
*/
func sortColors(nums []int) {
	n := len(nums)
	i, k := 0, n

	// get all k to the last
	for j := 0; j < k; j++ {
		if nums[j] == 2 {
			k--
			nums[k], nums[j] = nums[j], nums[k]
			j--
		}
	}
	// sort 0 and 1
	for j := 0; j < k; j++ {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
