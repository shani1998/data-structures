package tree

import (
	"fmt"
	"github.com/shani1998/data-structures/queue"
)

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	PreOrderTraversal(root.left)
	PreOrderTraversal(root.right)
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	PostOrderTraversal(root.left)
	PostOrderTraversal(root.right)
	fmt.Printf("%v ", root.val)
}

func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	InOrderTraversal(root.left)
	fmt.Printf("%v ", root.val)
	InOrderTraversal(root.right)
}

func LevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	// implement BFS algorithm using Queue
	queue := queue.NewQueue(100)
	queue.Push(root)

	for queue.Length() > 0 {
		item, _ := queue.Pop()
		fmt.Printf("%v-->", item.(*Node).val)
		if item.(*Node).left != nil {
			queue.Push(item.(*Node).left)
		}
		if item.(*Node).right != nil {
			queue.Push(item.(*Node).right)
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
			if item.(*Node).left != nil {
				queue.Push(item.(*Node).left)
			}
			if item.(*Node).right != nil {
				queue.Push(item.(*Node).right)
			}
		}

		// app levels to the result list
		result = append(result, levels)
	}

	return result

}
