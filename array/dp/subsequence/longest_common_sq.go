package subsequence

/*
https://leetcode.com/problems/longest-common-subsequence/
Given two strings text1 and text2, return the length of their longest common subsequence. If there is no common subsequence, return 0.

A subsequence of a string is a new string generated from the original string with some characters (can be none) deleted without changing the relative order of the remaining characters.
 For example, "ace" is a subsequence of "abcde".
A common subsequence of two strings is a subsequence that is common to both strings.

Example 1:
Input: text1 = "abcde", text2 = "ace"
Output: 3
*/

func lcsRecurse(s1, s2 string, i, j int) int {
    if i>=len(s1) || j >= len(s2) {
        return 0
    }
    if s1[i] == s2[j] {
        return 1 + lcsRecurse(s1, s2, i+1, j+1)
    }
    return max(lcsRecurse(s1, s2, i+1, j), lcsRecurse(s1, s2, i, j+1))
 }

 // longestCommonSubsequence returns the length of the longest common subsequence of text1 and text2.
 // time complexity: O(2^(m+n)), space complexity: O(m+n)
func longestCommonSubsequence(text1 string, text2 string) int {
    return lcsRecurse(text1, text2, 0, 0)
}

// DP approach: time complexity: O(m*n), space complexity: O(m*n)
func longestCommonSubsequenceDp(text1 string, text2 string) int {
    m, n := len(text1), len(text2)

    // DP table with extra row & column for base cases
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {

            if text1[i-1] == text2[j-1] {
                // characters match → extend LCS by 1
                dp[i][j] = dp[i-1][j-1] + 1
            } else {
                // take the best of skipping one char from either string
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }

    return dp[m][n]
}