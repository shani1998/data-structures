package questions

import "github.com/shani1998/data-structures/linkedlist"

/*
Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
https://www.codingninjas.com/blog/2020/09/09/floyds-cycle-detection-algorithm/
*/

// DetectCycle detects whether a cycle exist or not in the list
// if it exists then returns the starting point of the cycle.
func DetectCycle(head *linkedlist.SinglyNode) *linkedlist.SinglyNode {
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
