package main

import (
	"fmt"

	"github.com/AppBlitz/strutureData/list"
)

func main() {
	singlyLinkedList := &list.SinglyLinkedList{}
	singlyLinkedList.Append(2)
	singlyLinkedList.Append(4)
	singlyLinkedList.Append(7)
	// singlyLinkedList.RemoveAll()
	fmt.Printf("%d\n", singlyLinkedList.Len())
	fmt.Println(singlyLinkedList.IsEmpty())
}
