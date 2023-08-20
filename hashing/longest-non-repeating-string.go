package hashing

// lengthOfLongestSubstring returns the length of the longest substring without repeating characters, https://leetcode.com/problems/longest-substring-without-repeating-characters/

// naive approach time O(n^2) space O(n)
func lengthOfLongestSubstring(s string) int {
	n, maxLen := len(s), 0
	for i := 0; i < n; i++ {
		lookup := map[byte]struct{}{s[i]: {}}
		lenMap := 1
		for j := i + 1; j < n; j++ {
			if _, ok := lookup[s[j]]; ok {
				break
			}
			lookup[s[j]] = struct{}{}
			lenMap++
		}
		if lenMap > maxLen {
			maxLen = lenMap
		}
	}
	return maxLen
}

// using sliding window technique time O(n) space O(n)
func lengthOfLongestSubstring2(s string) int {
	lenStr := len(s)
	if lenStr <= 1 {
		return lenStr
	}

	maxLen := -1
	start := -1
	OccurredSoFar := make(map[byte]int)

	for i := 0; i < lenStr; i++ {
		// if char found in hash, shift start of char index and delete it from the set
		if pos, ok := OccurredSoFar[s[i]]; ok && pos > start {
			start = pos
			delete(OccurredSoFar, s[i])
		}

		OccurredSoFar[s[i]] = i
		if i-start > maxLen {
			maxLen = i - start
		}
	}

	return maxLen
}
