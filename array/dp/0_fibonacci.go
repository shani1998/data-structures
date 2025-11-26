package dp

/*
https://leetcode.com/problems/fibonacci-number/description/

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, 
such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).

Example 1:
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
*/

// Recursive approach: time complexity: O(2^n), space complexity: O(n) for recursion stack
func fibRecurse(n int) int {
    if n==0 || n == 1 {
        return n
    }

    return fibRecurse(n-1) + fibRecurse(n-2)
}

// DP approach: time complexity: O(n), space complexity: O(n)
func fibDP(n int) int {
	if n==0 || n == 1 {
		return n
	}

	// dpArr[i] will hold the Fibonacci number F(i)
	dpArr := make([]int, n+1)
	dpArr[0] = 0
	dpArr[1] = 1

	for i := 2; i <= n; i++ {
		dpArr[i] = dpArr[i-1] + dpArr[i-2]
	}

	return dpArr[n]
}

