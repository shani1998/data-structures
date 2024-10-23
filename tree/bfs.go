package tree

import (
	"fmt"

	"github.com/shani1998/data-structures/queue"
)

// LevelOrderTraversal tree level order traversal is the bfs implementation of tree
func LevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	// implement BFS algorithm using Queue
	queue := queue.NewQueue(100)
	queue.Push(root)

	for queue.Length() > 0 {
		item, _ := queue.Pop()
		fmt.Printf("%v-->", item.(*Node).Val)
		if item.(*Node).Left != nil {
			queue.Push(item.(*Node).Left)
		}
		if item.(*Node).Right != nil {
			queue.Push(item.(*Node).Right)
		}
	}

}

func GetLevelOrderTraversal(root *Node) [][]any {
	result := make([][]any, 0)
	if root == nil {
		return result
	}

	queue := queue.NewQueue(100)
	queue.Push(root)

	for queue.Length() > 0 {
		Qsize := queue.Length()
		levels := make([]any, 0)

		// pop all elements of current level
		// and add its children to the queue
		for i := 0; i < Qsize; i++ {
			item, _ := queue.Pop()
			levels = append(levels, item)
			if item.(*Node).Left != nil {
				queue.Push(item.(*Node).Left)
			}
			if item.(*Node).Right != nil {
				queue.Push(item.(*Node).Right)
			}
		}

		// app levels to the result list
		result = append(result, levels)
	}

	return result

}
