package linkedlist

import "fmt"

// ReverseList reverse the Linked-list of given head and return new head.
// i/p: 2-->3-->4-->5-->6, o/p: 6-->5-->4-->3-->2
// time: O(n), space: O(1)
func ReverseList(head *SinglyNode) *SinglyNode {
	// if there is no node in the list
	if head == nil {
		return nil
	}
	var prv *SinglyNode
	prv, curr := nil, head
	for curr != nil {
		temp := curr.next
		curr.next = prv
		prv = curr
		curr = temp
	}
	return prv
}

// ReverseListRecursive reverse the Linked-list and return new head.
// i/p: 2-->3-->4-->5-->6, o/p: 6-->5-->4-->3-->2
// time: O(n), space: O(n) <-- stack call
func ReverseListRecursive(head *SinglyNode) *SinglyNode {
	if head == nil || head.next == nil {
		return head
	}

	node := ReverseListRecursive(head.next)
	head.next.next = head // assign current address in next.next
	head.next = nil       // break the link of current node

	return node
}

// TraverseReverse print the linked list in reverse order.
func TraverseReverse(head *SinglyNode) {
	if head == nil {
		return
	}
	TraverseReverse(head.next)
	fmt.Printf("%v-->", head.data)
}
