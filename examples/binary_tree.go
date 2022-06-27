package main

import "github.com/shani1998/data-structures/tree"

func main() {
	var bt tree.BinaryTree

	bt.Insert(tree.NewNode("Hi"))
	bt.Insert(tree.NewNode("How"))
	bt.Insert(tree.NewNode("are"))
	bt.Insert(tree.NewNode("you"))
	bt.Insert(tree.NewNode("?"))

}
