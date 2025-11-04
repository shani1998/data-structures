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

func longestSubstringRecursive(s string, start int, set map[byte]bool) int {
	// base condition
	if start >= len(s) {
		return 0
	}

	maxLength := 0
	for i := start; i < len(s); i++ {
		if set[s[i]] {
			break
		}
		set[s[i]] = true
		maxLength = max(maxLength, i-start+1) //3
	}

	set = make(map[byte]bool)
	return max(maxLength, longestSubstringRecursive(s, start+1, set))
}

func longestSubstring(s string) int {
	return longestSubstringRecursive(s, 0, make(map[byte]bool))
}

func longestSubStrWithoutRepeatingCh(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	left, maxLen := 0, 0
	windowCh := make(map[byte]bool)

	for right := 0; right < n; right++ {
		// If char already in window, shrink from left
		for windowCh[s[right]] {
			delete(windowCh, s[left])
			left++
		}

		// Add current char and update maxLen
		windowCh[s[right]] = true
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

/*
Given a string s and an integer k, find the length of the longest substring that contains at most k distinct characters.
Input:  s = "eceba", k = 2 .. aa , 1 .. abababab, 2
Output: 3
Explanation: The substring is "ece" with length 3.
*/
// time O(n) as each ch processed twice by left/right ch, space O(k) for map
func longestSubStrLenLtK(s string, k int) int {
	left, maxLen := 0, 0
	windowCh := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		windowCh[s[right]]++

		// shrink window if distinct count > k
		for len(windowCh) > k {
			windowCh[s[left]]--
			if windowCh[s[left]] == 0 {
				delete(windowCh, s[left])
			}
			left++
		}

		// update max length when valid
		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
