package lists

type LinkedListCircular struct {
	Head  *Node
	Count int
}

func (listCircular *LinkedListCircular) AppendEnd(data int) {
	newNodes := NewNode(data)
	if listCircular.IsEmpty() {
		listCircular.Head = newNodes
		listCircular.Head.Next = listCircular.Head
	} else {
		currentNode := listCircular.Head
		for currentNode != listCircular.Head {
		}
	}
}

func (lislistCircular *LinkedListCircular) IsEmpty() bool {
	return lislistCircular.Head == nil
}
