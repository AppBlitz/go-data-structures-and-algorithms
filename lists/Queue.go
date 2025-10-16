package lists

var nodeEnd *Node

type Queue struct {
	Head *Node
}

func (queue *Queue) InQueue(data int) {
	newNodes := NewNode(data)
	if queue.Head == nil {
		queue.Head = newNodes
		nodeEnd = newNodes
	} else {
	}
}
