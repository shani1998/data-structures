package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var firstStr, lenStrs, result = strs[0], len(strs), ""
	for i, ch := range firstStr {
		for j := 1; j < lenStrs; j++ {
			if i >= len(strs[j]) || strs[j][i] != byte(ch) {
				return result
			}
		}
		result += string(ch)
	}
	return result
}
