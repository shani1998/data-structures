package slidingwindow

/*
Problem Statement

You are given an array of integers Arr and an integer k.
Your task is to find the maximum segment (subarray) of the array such that the segment contains no more than k unique integers.
If there are multiple segments with the same maximum length, return the segment with the lowest starting index.

The function should return the 1-based indices of the start and end of the segment.
Ex:

arr=[1,2,3,4,5] , k = 5
o/p : [1,5]

arr=[1,2,1,2,3,4,5] , k = 2
o/p : [1,4]

*/

// maxSegment1 is a brute force solution
func maxSegment1(k int, Arr []int) []int {
	start, end := 0, 0
	maxsegemet := 0
	if len(Arr) == 1 || k == 1 {
		return []int{1, 1}
	}
	for left := 0; left < len(Arr); left++ {
		lookup := map[int]struct{}{
			Arr[left]: {},
		}
		for right := left + 1; right < len(Arr); right++ {
			lookup[Arr[right]] = struct{}{}
			if len(lookup) <= k {
				maxsofar := right - left + 1
				if maxsofar > maxsegemet {
					maxsegemet = maxsofar
					start = left + 1
					end = right + 1
				}
			}
			if len(lookup) > k {
				break
			}
		}
	}
	return []int{start, end}
}

// maxSegment2 is a more optimized version using sliding window
// time complexity: O(n)
func maxSegment2(k int, Arr []int) []int {
	if len(Arr) == 0 || k == 0 {
		return []int{0, 0} // Edge case: No valid segment
	}

	lookup := make(map[int]int) // Map to store counts of elements in the window
	left, right := 0, 0         // Pointers for the sliding window
	maxSegment := 0             // Max segment length
	start, end := 0, 0          // Start and end indices of the result

	for right = 0; right < len(Arr); right++ {
		// Add the current element to the map
		lookup[Arr[right]]++

		// Shrink the window if unique elements exceed k
		for len(lookup) > k {
			lookup[Arr[left]]--
			if lookup[Arr[left]] == 0 {
				delete(lookup, Arr[left]) // Remove the element entirely
			}
			left++ // Move the left pointer
		}

		// Update maxSegment and indices if the current window is larger
		if right-left+1 > maxSegment {
			maxSegment = right - left + 1
			start = left + 1 // Convert to 1-based index
			end = right + 1  // Convert to 1-based index
		}
	}

	return []int{start, end}
}
