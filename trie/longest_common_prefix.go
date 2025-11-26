package trie

/*
https://leetcode.com/problems/longest-common-prefix/

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".
*/

// LongestCommonPrefix returns the longest prefix shared by every string in strs.
// Runs in O(n*m) where n is len(strs) and m is len(strs[0]).
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	lenFirstStr, lenStrs := len(first), len(strs)

	for i := 0; i < lenFirstStr; i++ {
		for j := 1; j < lenStrs; j++ {
			if i >= len(strs[j]) || strs[j][i] != first[i] {
				return first[:i]
			}
		}
	}

	return first
}

// LongestCommonPrefixWithTrie solves the problem using the trie defined in this package.
// time complexity: O(n*m), space complexity: O(n*m), where n is number of strings and m is average length of strings.
func LongestCommonPrefixWithTrie(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	trie := NewTrie()
	for _, str := range strs {
		if str == "" {
			return ""
		}
		trie.Insert(str)
	}

	return trie.LongestCommonPrefix()
}
