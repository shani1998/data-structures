package sorting

// SelectionSort algorithm sorts an array by repeatedly finding the minimum element
// from unsorted part and putting it at the beginning.
func SelectionSort(arr []int) []int {
	lenArr := len(arr)

	//  loop to access each array element
	for i := 0; i < lenArr; i++ {
		// find minimum element in unsorted part
		for j := i + 1; j < lenArr; j++ {
			// put min at the correct position
			if arr[j] < arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}

	return arr
}

/*
Worst and Average Case Time Complexity: O(N^2).
Best Case Time Complexity: O(N^2). The best case occurs when an array is already sorted.
Space Complexity: O(1)

-> Does sorting happen in place in selection sort?
Yes, it does not require extra space.

-> Is selection sort algorithm stable?
No because it changes the relative order of equal keys, but can be made : https://www.geeksforgeeks.org/stable-selection-sort/

*/
