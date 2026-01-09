package heap

/*
---------------------------------------------------------
 Divide Array in Groups of Size k
---------------------------------------------------------

Given an array and a number k, divide the array into groups of size k, considering the each groups are strictly increasing.
if it is not possible to divide the array in such groups, return false.

Example:
	Input: arr = [1,2,3,4,5,6], k = 3	
	Output: true
	Explanation: We can divide the array into two groups [1,2,3] and [4,5,6] which are strictly increasing.

	Input: arr = [1,2,3,2,3,5,7,8,9], k = 3
	Output: true
	Explanation: We can divide the array into three groups [1,2,3], [2,3,5] and [7,8,9] which are strictly increasing.
*/

func DivideArrayInGroups(arr []int, k int) bool {
	if len(arr)%k != 0 {
		return false
	}

	groups := len(arr) / k
	minHeap := &MinHeap{}
	heapInit(minHeap)

	for _, num := range arr {
		heapPush(minHeap, num)
	}

	for i := 0; i < groups; i++ {
		prev := -1
		for j := 0; j < k; j++ {
			if minHeap.size == 0 {
				return false
			}
			curr := heapPop(minHeap)
			if curr <= prev {
				return false
			}
			prev = curr
		}
	}

	return true
}