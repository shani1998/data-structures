package main

import (
	"fmt"
	"github.com/shani1998/data-structures/linkedlist"
)

func main() {
	doublyList := linkedlist.LinkedList{}
	for i := 0; i < 10; i++ {
		doublyList.InsertFirst(linkedlist.NewNode(i))
	}
	fmt.Println("Traversing forward")
	doublyList.TraverseForward()

	fmt.Println("Traversing backward")
	doublyList.TraverseBackward()

	fmt.Println("Len =", doublyList.Len())

	// insert at position 5
	err := doublyList.InsertAtIndex(5, linkedlist.NewNode(100))
	if err != nil {
		fmt.Printf("error inserting at index 3, %v\n", err)
	}
	fmt.Println("Traversing forward after inserting at index 5")
	doublyList.TraverseForward()

}
