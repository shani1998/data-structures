package recursion

/*
If your string is "abc",
then all possible subsequences (not substrings) are:

	           ""
	       /        \
	     "a"         ""
	   /    \       /    \
	 "ab"   "a"   "b"     ""
	/   \   / \   / \    / \

"abc" "ab" "ac" "a" "bc" "b" "c" ""
*/
func subsequences(s string) []string {
	var res []string
	var path []byte

	var backtrack func(i int)
	backtrack = func(i int) {
		// base case: reached end of string
		if i == len(s) {
			res = append(res, string(path))
			return
		}

		// ️Include current character
		path = append(path, s[i])
		backtrack(i + 1)

		// Exclude current character
		path = path[:len(path)-1] // backtrack
		backtrack(i + 1)
	}

	backtrack(0)
	return res
}
