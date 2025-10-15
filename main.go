package main

import (
	"github.com/AppBlitz/strutureData/list"
)

func main() {
	singlyLinkedList := &list.SinglyLinkedList{}
	singlyLinkedList.AppendEnd(2)
	singlyLinkedList.AppendEnd(4)
	singlyLinkedList.AppendStart(7)
	singlyLinkedList.AppendStart(5)
	singlyLinkedList.AppendStart(6)
	singlyLinkedList.AppendStart(9)
	singlyLinkedList.Display()
	// singlyLinkedList.RemoveAll()
	// fmt.Printf("%d\n", singlyLinkedList.Len())
	// fmt.Println(singlyLinkedList.IsEmpty())
	// fmt.Println(singlyLinkedList.GetValueNodePosition(2))
}
