package grid

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
	var result int
	dfs(0, 0, m, n, &result)
	return result
}

// dfs explores all valid paths by moving either down (r+1,c)
// or right (r,c+1). When it reaches the target cell, it counts 1 path.
// time complexity: O(2^(m+n)), space complexity: O(m+n)
func dfs(r, c, m, n int, result *int) {
	// Out of bounds → no further path
	if r >= m || c >= n {
		return
	}

	// Reached bottom-right → valid path found
	if r == m-1 && c == n-1 {
		*result++
		return
	}

	// Move DOWN
	dfs(r+1, c, m, n, result)

	// Move RIGHT
	dfs(r, c+1, m, n, result)
}

/*
The above naive DFS approach has an exponential time complexity O(2^(m+n)) because it explores all possible paths.
However, it has overlapping subproblems. The same cell can be reached via different paths,
leading to repeated calculations.
For example, consider a 3x3 grid:
                            (0,0)
                        /            \
                  (1,0)               (0,1)
                /      \            /        \
          (2,0)        (1,1)   (1,1)        (0,2)
           /   \       /   \    /   \       /     \
     (3,0)  (2,1)  (2,1) (1,2)(2,1)(1,2) (1,2)   (0,3)

This overlapping subproblem is evident as (1,1) and (2,1) are computed multiple times.
*/


// uniquePaths calculates the number of unique paths in an m x n grid.
// The robot can only move down or right. We use DFS + memoization (top-down DP) to avoid recomputing subproblems.
// time complexity: O(m*n), space complexity: O(m*n)
func uniquePathsDP(m int, n int) int {
    // Create DP table initialized with -1.
    // dp[r][c] will store the number of paths from cell (r,c) to the destination.
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = -1 // -1 means "not computed yet"
        }
    }

    // Start DFS from top-left corner (0,0)
    return dfsWithDP(0, 0, m, n, dp)
}

// dfs returns the number of unique paths from cell (r,c) to (m-1,n-1).
func dfsWithDP(r, c, m, n int, dp [][]int) int {
    // If out of bounds → no valid path
    if r < 0 || r >= m || c < 0 || c >= n {
        return 0
    }

    // If we reached the bottom-right cell → found 1 valid path
    if r == m-1 && c == n-1 {
        return 1
    }

    // If already computed, simply return stored result (memoization)
    if dp[r][c] != -1 {
        return dp[r][c]
    }

    // Otherwise compute:  paths by moving DOWN + paths by moving RIGHT
    dp[r][c] = dfsWithDP(r+1, c, m, n, dp) + dfsWithDP(r, c+1, m, n, dp)

    // Store and return the computed value
    return dp[r][c]
}