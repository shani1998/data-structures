package sorting

//BubbleSort it is a sorting algorithm that compares two adjacent elements
//and swaps them until they are not in the intended order.
func BubbleSort(arr []int) []int {
	//  loop to access each array element
	for i := 0; i < len(arr); i++ {
		swapped := false

		// loop to compare array elements
		for j := 0; j < len(arr)-i-1; j++ {
			// compare two adjacent elements
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		// no swapping means the array is already sorted
		// so no need for further comparison
		if !swapped {
			return arr
		}
	}
	return arr
}

/*
Worst and Average Case Time Complexity: O(N2).
Best Case Time Complexity: O(N). The best case occurs when an array is already sorted.
Space Complexity: O(1)

Does sorting happens in place in Bubble sort?
Yes, Bubble sort performs swapping of adjacent pairs without use of any major data structure. Hence Bubble sort algorithm is an in-place algorithm.

Is Bubble sort algorithm stable?
Yes, bubble sort algorithm is stable.

*/
