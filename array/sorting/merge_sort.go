package sorting

func mergeTwoSortedArrays(arr1, arr2 []int) []int {
	lenArr1, lenArr2 := len(arr1), len(arr2)
	result := make([]int, lenArr1+lenArr2)

	p, q, r := 0, 0, 0 // p, q, r are pointers to the array arr1, arr2 and result respectively

	// Until we reach either end of either arr1 or arr2, pick larger among
	// elements from arr1 and arr2 and place them in the correct position at result[p..r]
	for p < lenArr1 && q < lenArr2 {
		if arr1[p] < arr2[q] {
			result[r] = arr1[p]
			p++ // increment pointer to arr1 by 1
		} else {
			result[r] = arr2[q]
			q++ // increment pointer to arr2 by 1
		}
		r++
	}

	// When we run out of elements in either L or M,
	// pick up the remaining elements and put in A[p..r]
	if p < lenArr1 {
		for p < lenArr1 {
			result[r] = arr1[p]
			p++
			r++
		}
	}
	if q < lenArr2 {
		for q < lenArr2 {
			result[r] = arr2[q]
			r++
			q++
		}
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		// Dividing array into two half
		mid := len(arr) / 2
		// Sorting the first half
		left := MergeSort(arr[:mid])
		// Sorting the second half
		right := MergeSort(arr[mid:])
		return mergeTwoSortedArrays(left, right)
	}
	return arr
}
