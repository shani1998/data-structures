package tree

import "github.com/shani1998/data-structures/stack"

// tree dfs algorithms are inorder, preorder and postorder
// time complexity of the algorithms are O(n), visit all nodes exactly once
// while space complexity of these algorithms would be O(n) that is function call overhead or stack

/* InOrderTraversal Using Stack
1) Create an empty stack S.
2) Initialize current node as root
3) Push the current node to S and set current = current->left until current is NULL
4) If current is NULL and stack is not empty then
     a) Pop the top item from stack.
     b) Print the popped item, set current = popped_item->right
     c) Go to step 3.
5) If current is NULL and stack is empty then we are done.
*/

func InOrderTraversalIterative(root *Node) {
	var stack stack.Stack
	if root == nil {
		return
	}
	stack.Push(root)

	for len(stack) > 0 {
		curr := root.left
		// TODO
	}

}
