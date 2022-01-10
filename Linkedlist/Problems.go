package Linkedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// approach 1; space: O(n), Time:O(n)
// use stack: and multiply in last to first order with factor of 2
func getDecimalValue0(head *ListNode) int {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	// pop element from stack and multiply with pow(2,i)
	l, f, sum := len(stack), 1, 0
	for i := l - 1; i >= 0; i-- {
		sum = sum + stack[i]*f
		f = f * 2
	}
	return sum
}

// approach 2; space: O(1), Time:O(n)
// initialize the result as 0. Traverse the linked list and for each node,
// multiply the result by 2 and add the node’s data to it.
// (head.val)*2^2 + (head.val)*2^1 + (head.val)*2^0
func getDecimalValue1(head *ListNode) int {
	result := 0
	for head != nil {
		result = (result << 1) + head.Val
		head = head.Next
	}
	return result
}
