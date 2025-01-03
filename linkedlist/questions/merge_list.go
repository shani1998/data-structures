package questions

import "github.com/shani1998/data-structures/linkedlist"

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
func MergeTwoLists(list1, list2 *linkedlist.SinglyNode) *linkedlist.SinglyNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var head, curr *linkedlist.SinglyNode
	if list1.Val.(int64) <= list2.Val.(int64) {
		head, curr = list1, list1
		list1 = list1.Next
	} else {
		head, curr = list2, list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val.(int64) <= list2.Val.(int64) {
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
