package lists

import "fmt"

type Queue struct {
	Head  *Node
	Tail  *Node
	Count int
}

func (queue *Queue) EnQueue(data int) {
	newNode := NewNode(data)
	if queue.Tail == nil {
		queue.Head = newNode
		queue.Tail = newNode
		queue.Count = queue.Count + 1
	} else {
		queue.Count = 1 + queue.Count
		queue.Tail.Next = newNode
		queue.Tail = newNode
	}
}

func (queue *Queue) Dequeue() (int, bool) {
	if queue.Head == nil {
		return 0, false
	}
	removedData := queue.Head.Data
	queue.Head = queue.Head.Next
	if queue.Head == nil {
		queue.Tail = nil
	}

	queue.Count = queue.Count - 1
	return removedData, true
}

func (queue *Queue) Display() {
	currentNode := queue.Head
	for currentNode != nil {
		fmt.Printf("%d\n", currentNode.Data)
		currentNode = currentNode.Next
	}
}

func (queue *Queue) Len() int {
	return queue.Count
}
