package questions

import (
	"fmt"
	"github.com/shani1998/data-structures/tree"
)

func invertTree(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	queue := []*tree.Node{root}
	for len(queue) > 0 {
		item := queue[0]
		if item.Left != nil {
			queue = append(queue, item.Left)
		}
		if item.Right != nil {
			queue = append(queue, item.Right)
		}
		if len(queue) > 1 {
			queue = queue[1:]
		} else {
			queue = []*tree.Node{}
		}
		if item.Left != nil && item.Right != nil {
			item.Left, item.Right = item.Right, item.Left
		} else if item.Left != nil {
			item.Right = item.Left
			item.Left = nil
		} else {
			item.Left = item.Right
			item.Right = nil
		}
	}

	return root
}

func invertTree1(root *tree.Node) *tree.Node {
	if root == nil {
		return nil
	}

	queue := []*tree.Node{root}
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(len(queue))
		queue = queue[1:] // Dequeue the first element

		// Swap the left and right children
		node.Left, node.Right = node.Right, node.Left

		// Add the children to the queue if they are not nil
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}