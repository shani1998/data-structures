package slidingwindow

/*
 https://leetcode.com/problems/longest-repeating-character-replacement/description/

You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.
Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.
*/

// Brute force approach, time complexity O(n^2), space complexity O(n)
func characterReplacement(s string, k int) int {
	maxCount, lenS := 0, len(s)

	// AABABBA, 1
	for i := 0; i < lenS; i++ {
		freqMap := make(map[byte]int)
		maxFreqSoFar := 0

		for j := i; j < lenS; j++ {
			freqMap[s[j]]++
			maxFreqSoFar = max(maxFreqSoFar, freqMap[s[j]])
			windowSize := j - i + 1
			minChangesReq := windowSize - maxFreqSoFar

			// allowed changes
			if minChangesReq <= k {
				maxCount = max(maxCount, windowSize)
			} else {
				break
			}
		}
	}
	return maxCount
}

// using sliding window approach, time complexity O(n), space complexity O(n)
func characterReplacement1(s string, k int) int {
	freqMap := make(map[byte]int)
	left, maxCount, lenS := 0, 0, len(s)
	maxFreqSoFar := 0

	// ABAB, 2
	for right := 0; right < lenS; right++ {
		freqMap[s[right]]++

		maxFreqSoFar = max(maxFreqSoFar, freqMap[s[right]])

		windowSize := right - left + 1

		// is valid window with at most k changes
		if windowSize-maxFreqSoFar <= k {
			maxCount = max(maxCount, windowSize)
			continue
		}

		freqMap[s[left]]--
		left++
	}

	return maxCount
}
