package list

type Node struct {
	Data int
	Next *Node
}

type SinglyLinkedList struct {
	Head  *Node
	count int
}

func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

func (list *SinglyLinkedList) Append(Data int) {
	node := NewNode(Data)
	if list.Head == nil {
		list.Head = node
	} else {
		last := list.Head.Next
		for list.Head.Next != nil {
			last = list.Head.Next
		}
		last.Next = node
	}
}
