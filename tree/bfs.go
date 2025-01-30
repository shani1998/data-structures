package tree

import (
	"container/list"
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

// GetLevelOrderTraversal using channel as queue
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
			//node := queue[0]
			//queue = queue[1:]
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

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := list.New() // use container/list lib get O(1) time complexity for enqueue and dequeue operations, implemented using doubly linkedlist
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		levelVals := []int{}

		for i := 0; i < levelSize; i++ {
			elem := queue.Front() // get the first element
			node := elem.Value.(*Node)
			queue.Remove(elem)

			levelVals = append(levelVals, node.Val.(int))

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, levelVals)
	}

	return result
}

func levelOrderUsingArray(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*Node{root}
	front := 0 // Pointer to avoid slice shifting

	for front < len(queue) {
		levelSize := len(queue) - front
		levelVals := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++ // Move the front pointer instead of shifting elements

			levelVals = append(levelVals, node.Val.(int))

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelVals)
	}

	return result
}
