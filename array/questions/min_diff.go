package questions

import (
	"math"
	"sort"
)

/*
Given an array arr[] denoting heights of N towers and a positive integer K, you have to modify the height of each tower either by increasing or decreasing them by K only once.
Find out what could be the possible minimum difference of the height of shortest and longest towers after you have modified each tower.
Note: Assume that height of the tower can be negative.
A slight modification of the problem can be found here.
*/

// MinDiffHeightTowers find out what could be the possible minimum difference of the height of shortest and longest towers
func MinDiffHeightTowers(arr []int, n, k int) int {
	if n == 1 {
		return 0
	}

	sort.Ints(arr)
	ans := arr[n-1] - arr[0]

	for i := 1; i < n; i++ {
		if arr[i] >= k {
			minEle := int(math.Min(float64(arr[0]+k), float64(arr[i]-k)))
			maxEle := int(math.Max(float64(arr[n-1]-k), float64(arr[i-1]+k)))
			ans = int(math.Min(float64(ans), float64(maxEle-minEle)))
		}
	}

	return ans

	//sort.Ints(arr)
	//
	//minA, minB := arr[0]+k, arr[0]-k
	//maxA, maxB := arr[n-1]+k, arr[n-1]-k
	//
	//a := math.Abs(float64(minA - maxA))
	//b := math.Abs(float64(minA - maxB))
	//c := math.Abs(float64(minB - maxA))
	//d := math.Abs(float64(minB - maxB))
	//
	//return int(math.Min(math.Min(a, b), math.Min(c, d)))
}
