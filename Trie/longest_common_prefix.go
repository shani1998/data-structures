package Trie

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	srcStr, result := strs[0], ""
	for index, ch := range srcStr {
		notFound := false
		for i := 0; i < n; i++ {
			if index >= len(strs[i]) || string(strs[i][index]) != string(ch) {
				notFound = true
				break
			}
		}
		if notFound {
			break
		}
		result += string(ch)
	}
	return result

}
