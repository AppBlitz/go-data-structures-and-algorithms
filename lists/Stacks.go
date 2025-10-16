// Package lists implements structure stack of structure
package lists

import "fmt"

// Stack is of struct
type Stack struct {
	Head *Node
}

// Push add value to a stack
// Push(data int) void
func (stack *Stack) Push(data int) {
	currentNode := stack.Head
	if currentNode == nil {
		stack.Head = NewNode(data)
	} else {
		stack.Head = nil
		stack.Head = NewNode(data)
		stack.Head.Next = currentNode
	}
}

// Display show value of stack
// Display() void
func (stack *Stack) Display() {
	currentNode := stack.Head
	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}
}

// StackEmpty verification if stack is empty
// StackEmpty() bool
func (stack *Stack) StackEmpty() bool {
	return stack.Head == nil
}

// Pop implements  add data to a stack Pop () int
func (stack *Stack) Pop() int {
	currentNode := stack.Head
	removeNode := stack.Head.Next
	stack.Head = nil
	stack.Head = removeNode
	return currentNode.Data
}
