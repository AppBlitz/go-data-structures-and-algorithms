package lists

import (
	"errors"
	"fmt"
)

// Node structure, where it has the value to wait
// for and the reference to the next node
type Node struct {
	Data int
	Next *Node
}

// Type for create the singly linked list
type SinglyLinkedList struct {
	Head  *Node
	Count int
}

// Function for create a new node
func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

// function for add node in final list
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

// function return length of list
func (list *SinglyLinkedList) Len() int {
	return list.Count
}

// Function delete all elements of list
func (list *SinglyLinkedList) RemoveAll() {
	list.Head = nil
	list.Count = 0
}

// Function verification if list is empty
func (list *SinglyLinkedList) IsEmpty() bool {
	return list.Head == nil
}

// function show all elements of list
func (list *SinglyLinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

// Function for get data for position of list
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

// function to add a node to the beginning of the list;
// if the list is empty, it saves it as the header
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

// function for add node in a position
func (list *SinglyLinkedList) AppendNodeInPosition(data, positionNode int) {
	var nodePrevious *Node
	node := NewNode(data)
	currentNode := list.Head
	var auxiliaryPosition int
	for currentNode != nil {
		if auxiliaryPosition == positionNode {
			node.Next = currentNode
			nodePrevious.Next = nil
			nodePrevious.Next = node
			break
		}
		auxiliaryPosition = auxiliaryPosition + 1
		nodePrevious = currentNode
		currentNode = currentNode.Next
	}
}

func (list *SinglyLinkedList) RemoveElementPosition(positionElement int) {
	if positionElement == 0 {
		auxiliaryHead := list.Head.Next
		list.Head = nil
		list.Head = auxiliaryHead
	} else {
		currentNode := list.Head
		var axuliaryPreviousNode *Node
		var auxiliaryPosition int
		for currentNode != nil {
			if auxiliaryPosition == positionElement {
				node := currentNode.Next
				axuliaryPreviousNode.Next = nil
				axuliaryPreviousNode.Next = node
			}
			axuliaryPreviousNode = currentNode
			currentNode = currentNode.Next
			auxiliaryPosition = auxiliaryPosition + 1
		}
	}
}

func (list *SinglyLinkedList) SearchValueRecursive(data int) bool {
	return searchRecursivity(data, list.Head)
}

func searchRecursivity(data int, node *Node) bool {
	if node == nil {
		return false
	}
	if node.Data == data {
		return true
	}
	return searchRecursivity(data, node.Next)
}

func (list *SinglyLinkedList) AppendNodeEndListRecursive(data int) error {
	if !list.VerificationValueExist(data) {
		appendNodeEndListRecursive(data, list.Head)
	} else {
		return errors.New("Element exist")
	}
	return nil
}

func appendNodeEndListRecursive(data int, nodeList *Node) {
	if nodeList == nil {
		nodeList = NewNode(data)
	} else {
		appendNodeEndListRecursive(data, nodeList.Next)
	}
}
