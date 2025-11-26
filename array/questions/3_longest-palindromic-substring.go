package questions

/*
https://leetcode.com/problems/longest-palindromic-substring/description

Given a string s, return the longest palindromic substring in s.
Example 1:
Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
*/

// helper function to check if a string is palindrome
func isPalindrom(s string) bool {
    left, right := 0, len(s)-1
    for left < right {
        if s[left] != s[right] {
            return false
        }
        left++
        right--
    }
    return true
}

// Brute force approach, generate all possible substrings and check if it is palindrome 
// Time complexity: O(n^3), Space complexity: O(1)
func longestPalindrome(s string) string {
    n := len(s)
    if n <= 1 {
        return s
    }

    maxLen := 1
    result := s[0:1]

    for i := 0; i < n; i++ {
        for j := i + 1; j <= n; j++ { // must be <= n
            sub := s[i:j]
            if len(sub) > maxLen && isPalindrom(sub) {
                maxLen = len(sub)
                result = sub
            }
        }
    }
    return result
}

