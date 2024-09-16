package slidingwindow

/*
Given a string s, find the length of the longest substring without repeating characters.
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

func allUnique(s string, start int, end int) bool {
	lookup := make(map[byte]bool)
	for i := start; i < end; i++ {
		if lookup[s[i]] {
			return false
		}
		lookup[s[i]] = true
	}
	return true
}

// Brute force approach, Time complexity: O(n^3), Space complexity: O(min(n, m)),
// generate all possible substrings and check if it has unique characters
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	lookup := make(map[byte]int)
	maxLen, currMax := 0, 0
	i, n := 0, len(s)
	for i < n {
		indexAt, isExist := lookup[s[i]]
		if !isExist {
			lookup[s[i]] = i
			currMax++
			i++
			if currMax > maxLen {
				maxLen = currMax
			}
			continue
		}
		currMax = 0
		clear(lookup)
		i = indexAt + 1
	}
	return maxLen
}

// using sliding window approach, Time complexity: O(n), Space complexity: O(min(n, m)),
func lengthOfLongestSubstring3(s string) int {
	var left, right, maxLen int
	lookup := make(map[byte]int)
	lenStr := len(s)

	for right < lenStr {
		indexAt, isExist := lookup[s[right]]
		if isExist {
			if indexAt+1 > left { // check if the last occurrence of the character is within the current window
				left = indexAt + 1 // move left pointer to the right of the last occurrence of the character
			}
		}
		lenSubStr := right - left + 1
		maxLen = max(maxLen, lenSubStr)
		lookup[s[right]] = right
		right++
	}
	return maxLen
}
