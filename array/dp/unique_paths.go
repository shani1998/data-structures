package dp

// https://leetcode.com/problems/unique-paths/description/

/*

There is a robot on an m x n grid. The robot is initially located in the top-left corner (i.e., grid[0][0]).
The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.
Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

Input: m = 3, n = 7
Output: 28

Input: m = 3, n = 2
Output: 3

*/

func uniquePaths(m int, n int) int {
	return uniquePathsRecursive(m, n, 0, 0)
}

func uniquePathsRecursive(m, n, x, y int) int {
	// Base case: if we reach the bottom-right corner, return 1 (valid path)
	if x == m-1 && y == n-1 {
		return 1
	}

	// If out of bounds, return 0 (invalid path)
	if x >= m || y >= n {
		return 0
	}

	return uniquePathsRecursive(m, n, x+1, y) + uniquePathsRecursive(m, n, x, y+1)
}
