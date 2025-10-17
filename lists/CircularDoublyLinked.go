package lists

import "fmt"

type CircularDoublyLinke struct {
	Head  *Node
	Tail  *Node
	Count int
}

func (circularDoublyLinked *CircularDoublyLinke) RemoveAll() {
	if circularDoublyLinked.Head == nil {
		return
	} else {
		currentNode := circularDoublyLinked.Head
		for {
			next := currentNode.Next
			currentNode.Previous = nil
			currentNode.Next = nil
			if next == circularDoublyLinked.Head {
				break
			}
			currentNode = currentNode.Next
		}
		circularDoublyLinked.Head = nil
		circularDoublyLinked.Tail = nil
		circularDoublyLinked.Count = 0
	}
}

func (circularDoublyLinked *CircularDoublyLinke) IsEmpty() bool {
	return circularDoublyLinked.Head == nil || circularDoublyLinked.Tail == nil || circularDoublyLinked.Count == 0
}

func (circularDoublyLinked CircularDoublyLinke) Len() int {
	return circularDoublyLinked.Count
}

func (circularDoublyLinked *CircularDoublyLinke) AppendEnd(data int) {
	newNode := NewNode(data)

	if circularDoublyLinked.IsEmpty() {
		circularDoublyLinked.Head = newNode
		circularDoublyLinked.Tail = newNode
		newNode.Next = newNode
		newNode.Previous = newNode
		circularDoublyLinked.Count++
		return
	}
	newNode.Previous = circularDoublyLinked.Tail
	newNode.Next = circularDoublyLinked.Head
	circularDoublyLinked.Tail.Next = newNode
	circularDoublyLinked.Head.Previous = newNode
	circularDoublyLinked.Tail = newNode
	circularDoublyLinked.Count++
}

func (circularDoublyLinked *CircularDoublyLinke) AppendStart(data int) {
	newNode := NewNode(data)

	if circularDoublyLinked.IsEmpty() {
		circularDoublyLinked.Head = newNode
		circularDoublyLinked.Tail = newNode
		newNode.Next = newNode
		newNode.Previous = newNode
		circularDoublyLinked.Count++
		return
	}

	newNode.Next = circularDoublyLinked.Head
	newNode.Previous = circularDoublyLinked.Tail

	circularDoublyLinked.Tail.Next = newNode
	circularDoublyLinked.Head.Previous = newNode

	circularDoublyLinked.Head = newNode
	circularDoublyLinked.Count++
}

func (circularDoublyLinked *CircularDoublyLinke) AppendPosition(position int, data int) {
	newNode := NewNode(data)

	if position < 0 || position > circularDoublyLinked.Count {
		fmt.Println("Poisition invalid")
		return
	}

	if circularDoublyLinked.IsEmpty() || position == 0 {
		circularDoublyLinked.AppendStart(data)
		return
	}

	if position == circularDoublyLinked.Count {
		circularDoublyLinked.AppendEnd(data)
		return
	}

	current := circularDoublyLinked.Head
	for i := 0; i < position-1; i++ {
		current = current.Next
	}

	newNode.Next = current.Next
	newNode.Previous = current
	current.Next.Previous = newNode
	current.Next = newNode
	circularDoublyLinked.Count++
}

func (circularDoublyLinked *CircularDoublyLinke) GetValueAt(position int) int {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("list is empty")
		return -1
	}

	if position < 0 || position >= circularDoublyLinked.Count {
		fmt.Println("Posición inválida")
		return -1
	}

	current := circularDoublyLinked.Head
	for range position {
		current = current.Next
	}

	return current.Data
}

func (circularDoublyLinked *CircularDoublyLinke) Search(value int) bool {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("List empty")
		return false
	}

	current := circularDoublyLinked.Head
	for {
		if current.Data == value {
			return true
		}

		current = current.Next

		if current == circularDoublyLinked.Head {
			break
		}
	}

	return false
}

func (circularDoublyLinked *CircularDoublyLinke) GetPosition(value int) int {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("list is empty")
		return -1
	}

	current := circularDoublyLinked.Head
	position := 0

	for {
		if current.Data == value {
			return position
		}
		current = current.Next
		position++
		if current == circularDoublyLinked.Head {
			break
		}
	}
	return -1
}

func (circularDoublyLinked *CircularDoublyLinke) RemoveAt(position int) {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("Lista vacía")
		return
	}

	if position < 0 || position >= circularDoublyLinked.Count {
		fmt.Println("Posición inválida")
		return
	}

	if circularDoublyLinked.Count == 1 {
		circularDoublyLinked.Head = nil
		circularDoublyLinked.Tail = nil
		circularDoublyLinked.Count = 0
		return
	}

	if position == 0 {
		circularDoublyLinked.Head = circularDoublyLinked.Head.Next
		circularDoublyLinked.Head.Previous = circularDoublyLinked.Tail
		circularDoublyLinked.Tail.Next = circularDoublyLinked.Head
		circularDoublyLinked.Count--
		return
	}

	current := circularDoublyLinked.Head
	for range position {
		current = current.Next
	}

	current.Previous.Next = current.Next
	current.Next.Previous = current.Previous

	if current == circularDoublyLinked.Tail {
		circularDoublyLinked.Tail = current.Previous
	}

	circularDoublyLinked.Count--
}

func (circularDoublyLinked *CircularDoublyLinke) Display() {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("list is empty")
		return
	}

	current := circularDoublyLinked.Head
	for {
		fmt.Printf("%d ", current.Data)
		current = current.Next
		if current == circularDoublyLinked.Head {
			break
		}
	}
	fmt.Println()
}

func (circularDoublyLinked *CircularDoublyLinke) InsertarFinalRecursivo(valor int) {
	newNode := NewNode(valor)

	if circularDoublyLinked.IsEmpty() {
		circularDoublyLinked.Head = newNode
		circularDoublyLinked.Tail = newNode
		newNode.Next = newNode
		newNode.Previous = newNode
		circularDoublyLinked.Count++
		return
	}

	var insert func(current *Node)
	insert = func(current *Node) {
		if current.Next == circularDoublyLinked.Head {
			newNode.Previous = current
			newNode.Next = circularDoublyLinked.Head
			current.Next = newNode
			circularDoublyLinked.Head.Previous = newNode
			circularDoublyLinked.Tail = newNode
			circularDoublyLinked.Count++
			return
		}
		insert(current.Next)
	}
	insert(circularDoublyLinked.Head)
}

func (circularDoublyLinked *CircularDoublyLinke) InsertarInicioRecursivo(valor int) {
	newNode := NewNode(valor)

	if circularDoublyLinked.IsEmpty() {
		circularDoublyLinked.Head = newNode
		circularDoublyLinked.Tail = newNode
		newNode.Next = newNode
		newNode.Previous = newNode
		circularDoublyLinked.Count++
		return
	}

	var insert func(current *Node)
	insert = func(current *Node) {
		if current.Next == circularDoublyLinked.Head {
			newNode.Next = circularDoublyLinked.Head
			newNode.Previous = circularDoublyLinked.Tail
			circularDoublyLinked.Tail.Next = newNode
			circularDoublyLinked.Head.Previous = newNode
			circularDoublyLinked.Head = newNode
			circularDoublyLinked.Count++
			return
		}
		insert(current.Next)
	}
	insert(circularDoublyLinked.Head)
}

func (circularDoublyLinked *CircularDoublyLinke) InsertarEnPosicionRecursivo(posicion int, valor int) {
	if posicion < 0 || posicion > circularDoublyLinked.Count {
		fmt.Println("Position invalid")
		return
	}

	if circularDoublyLinked.IsEmpty() || posicion == 0 {
		circularDoublyLinked.InsertarInicioRecursivo(valor)
		return
	}

	if posicion == circularDoublyLinked.Count {
		circularDoublyLinked.InsertarFinalRecursivo(valor)
		return
	}

	newNode := NewNode(valor)

	var insert func(current *Node, index int)
	insert = func(current *Node, index int) {
		if index == posicion-1 {
			newNode.Next = current.Next
			newNode.Previous = current
			current.Next.Previous = newNode
			current.Next = newNode
			circularDoublyLinked.Count++
			return
		}
		insert(current.Next, index+1)
	}

	insert(circularDoublyLinked.Head, 0)
}

func (circularDoublyLinked *CircularDoublyLinke) BuscarRecursivo(valor int) bool {
	if circularDoublyLinked.IsEmpty() {
		fmt.Println("La lista está vacía")
		return false
	}

	var search func(current *Node) bool
	search = func(current *Node) bool {
		if current.Data == valor {
			return true
		}
		if current.Next == circularDoublyLinked.Head {
			return false
		}
		return search(current.Next)
	}

	return search(circularDoublyLinked.Head)
}
