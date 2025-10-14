package list

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head  *Node
	count int
}
