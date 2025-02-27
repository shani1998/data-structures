package backtracking

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
Given a pattern containing I for increasing and D for decreasing, you need to find the smallest number following that pattern.
You are given a 0-indexed string pattern of length n consisting of the characters 'I' meaning increasing and 'D' meaning decreasing.

A 0-indexed string num of length n + 1 is created using the following conditions:

num consists of the digits '1' to '9', where each digit is used at most once.
If pattern[i] == 'I', then num[i] < num[i + 1].
If pattern[i] == 'D', then num[i] > num[i + 1].
Return the lexicographically smallest possible string num that meets the conditions.

Example 1:

Input: pattern = "IIIDIDDD"
Output: "123549876"
Explanation:
At indices 0, 1, 2, and 4 we must have that num[i] < num[i+1].
At indices 3, 5, 6, and 7 we must have that num[i] > num[i+1].
Some possible values of num are "245639871", "135749862", and "123849765".
It can be proven that "123549876" is the smallest possible num that meets the conditions.
Note that "123414321" is not possible because the digit '1' is used more than once.
*/
func smallestNumber(pattern string) string {
	comb := smallestNumberUtil(0, pattern, "", math.MaxInt)
	return fmt.Sprintf("%d", comb)
}

func smallestNumberUtil(start int, pattern, result string, minNum int) int {
	fmt.Println(start, result, pattern, len(pattern), minNum)

	if start == len(pattern)+1 {
		num, _ := strconv.Atoi(result) // Convert string to int
		minNum = min(minNum, num)
		return minNum
	}

	for i := 1; i <= 9; i++ {
		if strings.Contains(result, fmt.Sprintf("%d", i)) {
			continue
		}
		if start > 0 && pattern[start-1] == 'I' && int(result[start-1]-'0') > i {
			continue
		}
		if start > 0 && pattern[start-1] == 'D' && int(result[start-1]-'0') < i {
			continue
		}
		minNum = smallestNumberUtil(start+1, pattern, result+fmt.Sprintf("%d", i), minNum)
	}

	return minNum
}
