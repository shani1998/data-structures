package linkedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Problem1:
	Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1.
	The linked list holds the binary representation of a number. Return the decimal value of the number in the linked list.
*/
func getDecimalValue(head *ListNode) int {
	/*
		// approach 1; space: O(n), Time:O(n)
		// use stack: and multiply in last to first order with factor of 2

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
	*/

	// approach 2; space: O(1), Time:O(n)
	// initialize the result as 0. Traverse the linked list and for each node,
	// multiply the result by 2 and add the node’s data to it.
	// (head.val)*2^2 + (head.val)*2^1 + (head.val)*2^0
	result := 0
	for head != nil {
		result = (result << 1) + head.Val
		head = head.Next
	}
	return result

}

/*
Problem2:
	Given the head of a singly linked list, return the middle node of the linked list.
	If there are two middle nodes, return the second middle node.
*/
func middleNode(head *ListNode) *ListNode {
	// Traverse linked list using two pointers. Move one pointer by one and the other pointers by two.
	// When the fast pointer reaches the end slow pointer will reach the middle of the linked list.
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // jump one step
		fast = fast.Next.Next // jump two-step
	}
	return slow
}

/*
Problem3:
	Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list,
	instead you will be given access to the node to be deleted directly.
	It is guaranteed that the node to be deleted is not a tail node in the list.
*/
func deleteNode(node *ListNode) {
	// It would be a simple deletion problem from the singly linked list if the head pointer was given
	// because for deletion you must know the previous node, and you can easily reach there by traversing from the head
	// pointer. Conventional deletion is impossible without knowledge of the prv node of a node that needs to be deleted.
	// The trick here is we can copy the data of the next node to the data field of the current node to be deleted
	// Then we can move one step forward. Now our next has become the current node and the current has become the
	// previous node. Now we can easily delete the current node by conventional deletion methods.
	temp := node.Next
	node.Val = temp.Val
	node.Next = temp.Next
	temp.Next = nil
}
