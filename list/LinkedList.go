package list

import "fmt"

type Data struct {
	Data  int
	Index int
}
type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head  *Node
	Count int
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func (list *SinglyLinkedList) Append(Data int) {
	node := NewNode(Data)
	if list.Head == nil {
		list.Head = node
		list.Count += 1
	} else {
		last := list.Head
		for last.Next != nil {
			last = last.Next
		}
		last.Next = node
		list.Count += 1
	}
}

func (list *SinglyLinkedList) Len() int {
	return list.Count
}

func (list *SinglyLinkedList) RemoveAll() {
	list.Head = nil
	list.Count = 0
}

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *SinglyLinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}
