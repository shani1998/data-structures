package backtracking

import "sort"

/*
A happy string is a string that:

consists only of letters of the set ['a', 'b', 'c'].
s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).
For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings and strings "aa", "baa" and "ababbc" are not happy strings.

Given two integers n and k, consider a list of all happy strings of length n sorted in lexicographical order.

Return the kth string of this list or return an empty string if there are less than k happy strings of length n.



Example 1:

Input: n = 1, k = 3
Output: "c"
Explanation: The list ["a", "b", "c"] contains all happy strings of length 1. The third string is "c".
*/

func getHappyString(n int, k int) string {
	happyStrings := getHappyStringUtil(0, n, "", []string{})
	if len(happyStrings) < k {
		return ""
	}
	sort.Strings(happyStrings)
	return happyStrings[k-1]
}

func getHappyStringUtil(start, n int, result string, combinations []string) []string {
	if len(result) == n {
		combinations = append(combinations, result)
		return combinations
	}

	for _, ch := range []string{"a", "b", "c"} {
		if start > 0 && string(result[start-1]) == ch {
			continue
		}
		combinations = getHappyStringUtil(start+1, n, result+ch, combinations)
	}

	return combinations
}

func getHappyString1(n int, k int) string {
	count := 0
	result := ""
	generateHappyString(0, n, "", k, &count, &result)
	return result
}

func generateHappyString(start, n int, current string, k int, count *int, result *string) {
	if len(current) == n {
		*count++
		if *count == k {
			*result = current
		}
		return
	}

	for _, ch := range []byte{'a', 'b', 'c'} {
		if start > 0 && current[start-1] == ch {
			continue
		}
		generateHappyString(start+1, n, current+string(ch), k, count, result)
		if *result != "" { // Stop early when k-th happy string is found
			return
		}
	}
}
