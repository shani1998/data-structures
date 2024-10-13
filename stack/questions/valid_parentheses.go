package questions

import (
	"github.com/shani1998/data-structures/stack"
)

/* Question: https://leetcode.com/problems/valid-parentheses/description/
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
*/

func isValid(s string) bool {
	// If the length is odd, it can't be balanced.
	if len(s)%2 != 0 {
		return false
	}

	stackArr := stack.Stack{}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stackArr.Push(ch)
		case ')', '}', ']':
			item, err := stackArr.Pop()
			if err != nil || item != brackets[ch] {
				return false
			}
		}
	}

	// If the stack is empty, all brackets were matched.
	return stackArr.Length() == 0
}
