package sorting

// InsertionSort is a sorting algorithm that places an unsorted
// element at its suitable place in each iteration.
func InsertionSort(arr []int) []int {
	lenArr := len(arr)

	// traverse through array 1 to lenArr
	// assume first element is sorted
	for i := 1; i < lenArr; i++ {
		key := arr[i]

		// Move elements of arr[0..i-1], that are
		// greater than key, to one position ahead
		// of their current position
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		// Place key at after the element just smaller than it.
		arr[j+1] = key
	}

	return arr
}

/*
Worst and Average Case Time Complexity: O(N^2).
Best Case Time Complexity: O(N). The best case occurs when an array is already sorted.
Space Complexity: O(1)


-> Does sorting happens in place in insertion sort?
Yes

-> Is insertion sort algorithm stable?
Yes,

-> Insertion sort is used when number of elements is small.
It can also be useful when input array is almost sorted, only few elements are misplaced in complete big array.

*/
