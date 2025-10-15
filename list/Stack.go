package list

type Stack struct {
	Head *Node
}

func (stack *Stack) Push(data int) {
	if stack.Head == nil {
		stack.Head = NewNode(data)
	} else {
		auxiliaryNode := stack.Head
		stack.Head = nil
		stack.Head = auxiliaryNode
	}
}

func (stack *Stack) StackEmpty() bool {
	return stack.Head == nil
}

func (stack *Stack) ReturnTopObject() int {
	return stack.Head.Data
}

func (stack *Stack) RemoveTopObject() {
	currentNode := stack.Head
	auxiliaryNode := stack.Head.Next
	currentNode = nil
	currentNode = auxiliaryNode
}
