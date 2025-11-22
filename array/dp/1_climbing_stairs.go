package dp

// climbStairsRecurse counts how many ways to reach exactly step n starting from step i.
// time complexity: O(2^n), space complexity: O(n)
func climbStairsRecurse(current, target int, result *int) {
    // If we cross beyond the target step, this is not a valid way
    if current > target {
        return
    }

    // If we reach exactly step target, one valid way is found
    if current == target {
        *result++
        return
    }

    // Try taking 1 step
    climbStairsRecurse(current+1, target, result)

    // Try taking 2 steps
    climbStairsRecurse(current+2, target, result)

}

// climbStairs returns the total number of distinct ways
// to reach step n when you can climb 1 or 2 steps at a time.
func climbStairs(n int) int {
    var result int
    climbStairsRecurse(0, n, &result)

    return result
}

/*
The above naive recursive approach has an exponential time complexity O(2^n) because it explores all possible ways.
However, it has overlapping subproblems. The same step can be reached via different paths,
leading to repeated calculations.
For example, consider n = 4:
						
                                       f(4)
                                /               \
                           f(3)                 f(2)
                         /      \             /      \
                    f(2)        f(1)      f(1)      f(0)
                   /    \
              f(1)     f(0)
		Here, f(2) and f(1) are computed multiple times.
*/


// This uses DFS + memoization (top-down DP). time complexity: O(n), space complexity: O(n)
func climbStairsDP1(n int) int {
    // dp[i] will store the number of ways to reach step n starting from step i.
    dp := make([]int, n+1)

    // Initialize dp array with -1 meaning "not computed yet"
    for i := 0; i <= n; i++ {
        dp[i] = -1
    }

    return countWaysDPUtil(0, n, dp)
}

// countWays returns the number of ways to go from step 'current' to 'target'.
func countWaysDPUtil(current, target int, dp []int) int {
    // If we overshoot the target step, it's not a valid way.
    if current > target {
        return 0
    }

    // If we land exactly on the target step, we found one valid way.
    if current == target {
        return 1
    }

    // If already computed, return memoized result.
    if dp[current] != -1 {
        return dp[current]
    }

    // Otherwise compute:
    // ways = ways by taking 1 step + ways by taking 2 steps
    dp[current] = countWaysDPUtil(current+1, target, dp) +
                  countWaysDPUtil(current+2, target, dp)

    return dp[current]
}

// climbStairsDP2 uses bottom-up DP approach. time complexity: O(n), space complexity: O(n)
func climbStairsDP2(n int) int {
	// if we see the output pattern,
	// n = 1 o/p: 1
	// n = 2 o/p: 2
	// n = 3 o/p: 3 ie op1 + op2
	// n = 4 o/p: 5 ie op3 + op2
	// ... same as fibonacci number
	if n < 2 {
		return n
	}
	dpArr := make([]int, n)
	dpArr[0] = 1
	dpArr[1] = 2
	for i := 2; i < n; i++ {
		dpArr[i] = dpArr[i-1] + dpArr[i-2]
	}
	return dpArr[n-1]
}

/*
https://leetcode.com/problems/min-cost-climbing-stairs/
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.

Example 1:
Input: cost = [10,15,20]
Output: 15
Explanation: You will start at index 1.
- Pay 15 and climb two steps to reach the top.
The total cost is 15.
*/

// Recursive function to compute minimum cost starting from step i
func minCostClimbingStairsRecurse(i, n int, cost []int) int {
    // If we go past the last step, no additional cost
    if i > n {
        return 0
    }

    // If we land exactly on the last step, pay its cost
    if i == n {
        return cost[i]
    }

    // Choose minimum cost between taking 1 step or 2 steps forward
    return min(
        cost[i] + minCostClimbingStairsRecurse(i+1, n, cost),
        cost[i] + minCostClimbingStairsRecurse(i+2, n, cost),
    )
}

func minCostClimbingStairs(cost []int) int {
    // You can start from step 0 or step 1, take the cheaper path
    return min(
        minCostClimbingStairsRecurse(0, len(cost)-1, cost),
        minCostClimbingStairsRecurse(1, len(cost)-1, cost),
    )
}

func minCostClimbingStairsDPUtil(i, n int, cost []int, dp []int) int {
    // If we go past the last real step, no cost
    if i > n {
        return 0
    }

    // If we're on the last valid cost index
    if i == n {
        return cost[i]
    }

    // Memoized
    if dp[i] != -1 {
        return dp[i]
    }

    dp[i] = cost[i] + min(
        minCostClimbingStairsDPUtil(i+1, n, cost, dp),
        minCostClimbingStairsDPUtil(i+2, n, cost, dp),
    )

    return dp[i]
}

func minCostClimbingStairsDP(cost []int) int {
    n := len(cost) - 1  // last real step index
    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = -1
    }

    // Start from step 0 or step 1
    return min(
        minCostClimbingStairsDPUtil(0, n, cost, dp),
        minCostClimbingStairsDPUtil(1, n, cost, dp),
    )
}
