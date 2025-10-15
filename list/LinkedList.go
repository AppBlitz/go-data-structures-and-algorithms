package list

import (
	"errors"
	"fmt"
)

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

func (list *SinglyLinkedList) AppendEnd(Data int) {
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

func (list *SinglyLinkedList) GetValueNodePosition(positionNode int) (int, error) {
	var auxiliaryPosition int
	if list.Len() < positionNode {
		return -1, errors.New("the index not valid in the list ")
	} else {
		var counterIndexNode int
		currentNode := list.Head
		for currentNode != nil {
			if counterIndexNode == positionNode {
				auxiliaryPosition = currentNode.Data
			}
			counterIndexNode = counterIndexNode + 1
			currentNode = currentNode.Next
		}

	}
	return auxiliaryPosition, nil
}

// This function vericate if value exist in list
// receive data
func (list *SinglyLinkedList) VerificationValueExist(data int) bool {
	currentNode := list.Head
	var auxiliaryExistValue bool
	for currentNode != nil {
		if currentNode.Data == data {
			auxiliaryExistValue = true
			break
		}
		currentNode = currentNode.Next
	}
	return auxiliaryExistValue
}

func (list *SinglyLinkedList) AppendStart(data int) {
	node := NewNode(data)
	if list.Head == nil {
		list.Head = node
	} else {
		nodeAuxHead := list.Head
		list.Head = node
		list.Head.Next = nodeAuxHead
	}
}
