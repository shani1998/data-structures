package searching

import "fmt"

// Python Program to search an element
// in a sorted and pivoted array
// Searches an element key in a pivoted
// sorted array arrp[] of size n
//
// https://www.geeksforgeeks.org/search-an-element-in-a-sorted-and-pivoted-array/
// Input  : arr[] = {5, 6, 7, 8, 9, 10, 1, 2, 3};
// key = 3
// Output : Found at index 8
// https://www.youtube.com/watch?v=oTfPJKGEHcc

//get the position of pivot element
func getPivotElementPosition(nums []int, low int, high int) int {
	if high < low {
		return -1
	}
	if high == low {
		return low
	}
	mid := int((low + high) / 2)

	if mid < high && nums[mid] > nums[mid+1] {
		return mid
	}

	if mid > low && nums[mid] < nums[mid-1] {
		return mid - 1
	}

	if nums[low] < nums[mid] {
		return getPivotElementPosition(nums, mid+1, high)
	}
	return getPivotElementPosition(nums, low, mid-1)
}

//idea is to find the pivot position and then partition array into two sorted array
//and search key into sorted arrays complexity O(2logn)
func search(nums []int, target int) int {
	pivot := getPivotElementPosition(nums, 0, len(nums)-1)
	fmt.Println(pivot)
	if pivot == -1 {
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if nums[pivot] == target {
		return pivot
	}

	//check if target falls in left half or right half of the pivot element
	if nums[0] <= target && nums[pivot] > target {
		return binarySearch(nums, 0, pivot, target)
	}
	return binarySearch(nums, pivot+1, len(nums)-1, target)

}

//----------------------------approach-2-----------------------------only O(logn)
func search2(nums []int, target int) int {
	var low, high = 0, len(nums) - 1
	for low <= high {
		mid := int((low + high) / 2)
		if nums[mid] == target {
			return mid
		}
		fmt.Println(low, high, mid)
		//if left half is uniformly increasing curve
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] { //check if target exist in left half or not
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[high] { //check if target exist in right half or not
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1

}
