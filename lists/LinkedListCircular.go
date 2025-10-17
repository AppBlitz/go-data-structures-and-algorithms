package lists

import "fmt"

type LinkedListCircular struct {
	Head  *Node
	Count int
}

func (listCircular *LinkedListCircular) AppendEnd(data int) {
	newNodes := NewNode(data)
	if listCircular.IsEmpty() {
		listCircular.Head = newNodes
		listCircular.Head.Next = listCircular.Head
		listCircular.Count = listCircular.Count + 1
	} else {
		currentNode := listCircular.Head
		for currentNode != listCircular.Head {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNodes
		newNodes.Next = listCircular.Head
		listCircular.Count = listCircular.Count + 1
	}
}

func (listCircular *LinkedListCircular) AppendStart(data int) {
	newNodes := NewNode(data)
	if listCircular.IsEmpty() {
		listCircular.Head = newNodes
		listCircular.Head.Next = newNodes
		listCircular.Count = listCircular.Count + 1
	} else {
		newNodes.Next = listCircular.Head
		listCircular.Head = nil
		listCircular.Head = newNodes
	}
}

func (listCircular LinkedListCircular) Len() int {
	return listCircular.Count
}

func (listCircular *LinkedListCircular) RemoveAll() {
	listCircular.Head = nil
}

func (listCircular *LinkedListCircular) IsEmpty() bool {
	return listCircular.Head == nil
}

func (listCircular *LinkedListCircular) AppendPosition(data, position int) {
	newNode := NewNode(data)

	if listCircular.Head == nil {
		listCircular.Head = newNode
		newNode.Next = newNode
		listCircular.Count++
		return
	}

	if position == 0 {
		last := listCircular.Head
		for last.Next != listCircular.Head {
			last = last.Next
		}
		newNode.Next = listCircular.Head
		listCircular.Head = newNode
		last.Next = newNode
		listCircular.Count++
		return
	}

	current := listCircular.Head
	var prev *Node
	count := 0

	for count < position {
		prev = current
		current = current.Next
		count++

		if current == listCircular.Head {
			break
		}
	}

	newNode.Next = current
	prev.Next = newNode
	listCircular.Count++
}

func (listCircular *LinkedListCircular) GetValue(position int) int {
	var auxiliaryValue int
	var auxiliaryCount int
	currentNode := listCircular.Head
	for currentNode != listCircular.Head {
		if auxiliaryCount == position {
			auxiliaryValue = listCircular.Head.Data
		}
		auxiliaryCount = auxiliaryCount + 1
		currentNode = currentNode.Next
	}
	return auxiliaryValue
}

func (listCircular *LinkedListCircular) SearchValue(data int) bool {
	currentNode := listCircular.Head
	var auxiliaryValueExist bool
	for currentNode != listCircular.Head {
		if currentNode.Data == data {
			auxiliaryValueExist = true
		}
		currentNode = currentNode.Next
	}
	return auxiliaryValueExist
}

func (listCircular *LinkedListCircular) Display() {
	nodeCurrent := listCircular.Head
	for nodeCurrent != listCircular.Head {
		fmt.Println(listCircular.Head.Data)
		nodeCurrent = nodeCurrent.Next
	}
}

func (listCircular *LinkedListCircular) SearchValueRecursive(data int) bool {
	return searchValueRecursive(data, listCircular.Head, listCircular.Head.Next)
}

func searchValueRecursive(data int, auxiliaryNode *Node, nodeNext *Node) bool {
	if data == auxiliaryNode.Data {
		return true
	}
	if auxiliaryNode == nodeNext {
		return false
	}
	return searchValueRecursive(data, auxiliaryNode, nodeNext.Next)
}

func (listCircular *LinkedListCircular) GetPositionElement(data int) int {
	var auxiliaryIndex int
	var auxiliaryCount int
	currentNode := listCircular.Head
	for currentNode != listCircular.Head {
		if currentNode.Data == data {
			auxiliaryIndex = auxiliaryCount
		}
		auxiliaryCount = auxiliaryCount + 1
		currentNode = currentNode.Next
	}
	return auxiliaryIndex
}

func (listCircular *LinkedListCircular) InserEndList(valor int) {
	newNode := NewNode(valor)

	if listCircular.Head == nil {
		listCircular.Head = newNode
		newNode.Next = newNode
		listCircular.Count++
		return
	}

	var insertar func(current *Node)
	insertar = func(current *Node) {
		if current.Next == listCircular.Head {
			current.Next = newNode
			newNode.Next = listCircular.Head
			listCircular.Count++
			return
		}
		insertar(current.Next)
	}
	insertar(listCircular.Head)
}

func (listCircular *LinkedListCircular) InsertInpositionRecursive(posicion int, valor int) {
	if posicion < 0 || posicion > listCircular.Count {
		fmt.Println("Position Invalid")
		return
	}

	newNode := NewNode(valor)

	if listCircular.Head == nil {
		listCircular.Head = newNode
		newNode.Next = newNode
		listCircular.Count++
		return
	}

	if posicion == 0 {
		current := listCircular.Head
		for current.Next != listCircular.Head {
			current = current.Next
		}
		newNode.Next = listCircular.Head
		listCircular.Head = newNode
		current.Next = newNode
		listCircular.Count++
		return
	}

	var insertar func(current *Node, index int)
	insertar = func(current *Node, index int) {
		if index == posicion-1 {
			newNode.Next = current.Next
			current.Next = newNode
			listCircular.Count++
			return
		}
		insertar(current.Next, index+1)
	}
	insertar(listCircular.Head, 0)
}

func (listCircular *LinkedListCircular) RemovePoisition(posicion int) {
	if listCircular.Head == nil || posicion < 0 || posicion >= listCircular.Count {
		fmt.Println("Posición inválida o lista vacía")
		return
	}

	if posicion == 0 {
		// eliminar el primero
		if listCircular.Count == 1 {
			listCircular.Head = nil
		} else {
			current := listCircular.Head
			for current.Next != listCircular.Head {
				current = current.Next
			}
			listCircular.Head = listCircular.Head.Next
			current.Next = listCircular.Head
		}
		listCircular.Count--
		return
	}

	var eliminar func(current *Node, index int)
	eliminar = func(current *Node, index int) {
		if index == posicion-1 {
			current.Next = current.Next.Next
			listCircular.Count--
			return
		}
		eliminar(current.Next, index+1)
	}
	eliminar(listCircular.Head, 0)
}
