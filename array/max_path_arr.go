package main

import "fmt"

/*
Given two sorted arrays, the arrays may have some common elements. Find the sum of the maximum sum path to reach from the beginning of any array to the
end of any of the two arrays. We can switch from one array to another array only at common elements.

Note: The common elements do not have to be at the same indexes.

Examples: https://www.geeksforgeeks.org/maximum-sum-path-across-two-arrays/
Input: ar1[] = {2, 3, 7, 10, 12}, ar2[] = {1, 5, 7, 8}
Output: 35
Explanation: 35 is sum of 1 + 5 + 7 + 10 + 12.
Start from the first element of ar2 which is 1, then move to 5, then 7.  From 7 switch to ar1 (as 7 is common) and traverse 10 and 12.
*/
func maxSumPath(arr1, arr2 []int, m, n int) int {
	arr1Sum, arr2Sum := 0, 0
	maxSum := 0

	i, j := 0, 0
	for i < m && j < n {
		if arr1[i] == arr2[j] {
			maxSum += max(arr1Sum, arr2Sum) + arr1[i]
			arr1Sum, arr2Sum = 0, 0
			i++
			j++
			continue
		}
		if arr1[i] < arr2[j] {
			arr1Sum += arr1[i]
			i++
		} else {
			arr2Sum += arr2[j]
			fmt.Printf("arr1[%d]: %d, arr2[%d]: %d, arr1Sum: %d, arr2Sum: %d, maxSum: %d \n", i, arr1[i], j, arr2[j], arr1Sum, arr2Sum, maxSum)
			j++
		}
	}
	for i < m {
		arr1Sum += arr1[i]
		i++
	}
	for j < n {
		arr2Sum += arr2[j]
		j++
	}
	maxSum += max(arr1Sum, arr2Sum)

	return maxSum
}
