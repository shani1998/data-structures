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

/*
	MergeTwoLists
	You are given the heads of two sorted linked lists list1 and list2.
	Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
	Return the head of the merged linked list.
	Ex:
	list1: 1->3->5
	list2: 2->4->6
	o/p: 1->2->3->4->5
*/
// MergeTwoLists  takes two lists as input and returns
// head of the merged sorted list. time: O(len(list1)+len(list2))
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head, curr *ListNode
	if list1.Val <= list2.Val {
		head, curr = list1, list1
		list1 = list1.Next
	} else {
		head, curr = list2, list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	if list1 != nil {
		curr.Next = list1
	}

	if list2 != nil {
		curr.Next = list2
	}

	return head
}

/*
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
https://www.codingninjas.com/blog/2020/09/09/floyds-cycle-detection-algorithm/
*/

// DetectCycle detects whether a cycle exist or not in the list
// if it exists then returns the starting point of the cycle.
func DetectCycle(head *ListNode) *ListNode {
	// use floyd cycle detection algorithm

	// detect whether cycle exist or not
	var slow, fast = head, head
	var isCycleFound bool
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycleFound = true
			break
		}
	}

	// return nil if there is no cycle
	if !isCycleFound {
		return nil
	}

	// reset slow pointer to head
	// and jump both pointer one step at a time
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}