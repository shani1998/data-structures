package questions

import "github.com/shani1998/data-structures/linkedlist"

/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
The first node is considered odd, and the second node is even, and so on.
Note that the relative order inside both the even and odd groups should remain as it was in the input.

ex:
Input: head = [1,2,3,4,5]
Output: [1,3,5,2,4]
*/

func oddEvenList(head *linkedlist.SinglyNode) *linkedlist.SinglyNode {
	if head == nil {
		return head
	}

	var eventPtr, evenPtrHead *linkedlist.SinglyNode
	oddPtr, oddPtrHead, c := head, head, 1

	for head.Next != nil {
		head = head.Next
		c++
		if c%2 == 1 {
			oddPtr.Next = head
			oddPtr = oddPtr.Next
		} else {
			if eventPtr == nil {
				eventPtr = head
				evenPtrHead = eventPtr
			} else {
				eventPtr.Next = head
				eventPtr = eventPtr.Next
			}
		}
	}

	oddPtr.Next = evenPtrHead
	if eventPtr != nil {
		eventPtr.Next = nil
	}

	return oddPtrHead
}

/**
 * Definition for singly-linked list.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */
func oddEvenList1(head *linkedlist.SinglyNode) *linkedlist.SinglyNode {
	if head == nil {
		return head
	}

	oddPtr, evenPtr := head, head.Next // Odd and even pointers
	evenHead := evenPtr                // Save the head of the even list

	for oddPtr != nil && evenPtr != nil && evenPtr.Next != nil {
		oddPtr.Next = evenPtr.Next // Link the next odd node
		oddPtr = oddPtr.Next       // Move the odd pointer
		evenPtr.Next = oddPtr.Next // Link the next even node
		evenPtr = evenPtr.Next     // Move the even pointer
	}

	oddPtr.Next = evenHead // Connect the odd list to the even list
	return head            // Return the updated head
}
