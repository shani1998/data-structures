package dp

func climbStairs(n int) int {
	totalWays := 0
	climbStairsUtil(0, n, &totalWays)
	return totalWays
}

// bottom-up approach, start from 0 and go till n
func climbStairsUtil(stair, n int, totalWays *int) {
	if stair == n {
		*totalWays++
		return
	}
	if stair+1 <= n {
		climbStairsUtil(stair+1, n, totalWays)
	}
	if stair+2 <= n {
		climbStairsUtil(stair+2, n, totalWays)
	}
}

// top-down approach, start from n and go till 0
func climbStairsRecursive(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsDP(n int) int {
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
stair = 0
    +1 --> climbStairsUtil(1,0,3) <-- 2+1 climbStairsUtil(3,1,3) --> return 2
    +1 +1 --> climbStairsUtil(2,0,3) <-- 2+2 can't take two steps --> return
    +1 +1 +1 --> climbStairsUtil(3,0,3) --> 1

    +2 --> climbStairsUtil(2,0,3)
    +2 + 1 --> climbStairsUtil(3,0,3) = 3

                climbStairs(0)
               /          \
     climbStairs(1)       climbStairs(2)
     /       \                 /       \
climbStairs(2)  climbStairs(3) climbStairs(3) climbStairs(4)
    /     \                      (Base Case)      (Out of bounds)
climbStairs(3) climbStairs(4)
(Base Case)  (Out of bounds)

if you notice here climbStairs(2) , recomputed twice, we can avoid this using dp
*/
