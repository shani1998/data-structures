package backtracking

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
Ex1:
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]

Ex2:
Input: n = 1
Output: ["()"]
*/

// Approach 1 - generate all permutations of the parentheses string and validate each permutation
func GenerateParentheses(n int) []string {
	// Generate the initial parentheses string with n '(' and n ')'
	var parenthesesStr []rune
	for i := 0; i < n; i++ {
		parenthesesStr = append(parenthesesStr, '(', ')')
	}

	// Use a map to store unique permutations
	uniqueResults := make(map[string]bool)
	generateParenthesesUtil(parenthesesStr, 0, uniqueResults)

	// Convert the map to a slice of strings
	var finalResult []string
	for key := range uniqueResults {
		finalResult = append(finalResult, key)
	}

	return finalResult
}

func generateParenthesesUtil(permuteStr []rune, start int, result map[string]bool) {
	if start == len(permuteStr)-1 && isValidParentheses(permuteStr) && !result[string(permuteStr)] {
		result[string(permuteStr)] = true // Add to the set
		return
	}

	for i := start; i < len(permuteStr); i++ {
		// Swap the elements to generate a new permutation
		permuteStr[i], permuteStr[start] = permuteStr[start], permuteStr[i]
		generateParenthesesUtil(permuteStr, start+1, result)
		// Backtrack: Swap back to restore the original order
		permuteStr[i], permuteStr[start] = permuteStr[start], permuteStr[i]
	}

	return
}

func isValidParentheses(str []rune) bool {
	stack := []rune{}

	// Validate the parentheses using a stack
	for i := 0; i < len(str); i++ {
		if str[i] == ')' {
			// Check for a matching '('
			stackLen := len(stack)
			if stackLen-1 < 0 || stack[stackLen-1] != '(' {
				return false
			}
			// Pop the top element from the stack
			stack = stack[:stackLen-1]
		} else if str[i] == '(' {
			// Push '(' onto the stack
			stack = append(stack, '(')
		}
	}

	// Ensure the stack is empty for a valid parentheses sequence
	return len(stack) == 0
}
