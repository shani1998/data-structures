package searching

//Recursive Approach: O(logn)
func binarySearch[T int | float64 | string](arr []T, low int, high int, key T) int {
	// We reach here when element is not present in array
	if high < low {
		return -1
	}

	mid := (low + high) / 2

	// If the element is present at the middle
	if arr[mid] == key {
		return mid
	}

	// If element is smaller than mid, then  it can only be present in left subarray
	if arr[mid] > key {
		return binarySearch(arr, low, mid-1, key)
	}

	// Else the element can only be present in right subarray
	return binarySearch(arr, mid+1, high, key)
}

//Iterative approach: O(logn)
func binarySearchIterative[T int | float64 | string](arr []T, key T) int {
	var low, high = 0, len(arr) - 1
	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == key {
			return mid
		}
		if arr[mid] < key {
			//target might exist in right half
			low = mid + 1
		} else {
			//target might exist in left half
			high = mid - 1
		}
	}
	return -1

}
