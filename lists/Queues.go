package lists

import "fmt"

type Queue struct {
	Head *Node
	Tail *Node
}

func (queue *Queue) EnQueue(data int) {
	newNode := NewNode(data)
	if queue.Tail == nil {
		queue.Head = newNode
		queue.Tail = newNode
	} else {
		queue.Tail.Next = newNode
		queue.Tail = newNode
	}
}

func (queue *Queue) Dequeue() (int, bool) {
	if queue.Head == nil {
		return 0, false // Cola vacía
	}
	removedData := queue.Head.Data
	queue.Head = queue.Head.Next
	if queue.Head == nil {
		queue.Tail = nil // Si la cola queda vacía
	}
	return removedData, true
}

func (queue *Queue) Display() {
	currentNode := queue.Head
	for currentNode != nil {
		fmt.Printf("%d\n", currentNode.Data)
		currentNode = currentNode.Next
	}
}
