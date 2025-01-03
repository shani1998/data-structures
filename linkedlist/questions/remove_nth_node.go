package questions

import "github.com/shani1998/data-structures/linkedlist"

// RemoveNthFromEnd Given the head of a linked list, remove the nth node from the end of the list and return its head.
// problem link: https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
func RemoveNthFromEnd(head *linkedlist.SinglyNode, n int) *linkedlist.SinglyNode {
	totalNodes := getLenOfLinkList(head)
	desireNodePos := (totalNodes - n) + 1

	if desireNodePos <= 0 {
		return head
	}

	curr, prev, currNodePos := head, head, 1
	for currNodePos != desireNodePos {
		prev = curr
		curr = curr.Next
		currNodePos++
	}

	// if head node is desire node to be deleted
	if curr == prev {
		head = head.Next
		prev.Next = nil
		return head
	}

	prev.Next = curr.Next

	return head
}

func getLenOfLinkList(head *linkedlist.SinglyNode) int {
	if head == nil {
		return 0
	}
	i, temp := 1, head
	for temp.Next != nil {
		i++
		temp = temp.Next
	}
	return i
}
