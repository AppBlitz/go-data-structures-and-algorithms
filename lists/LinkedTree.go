package lists

import "fmt"

// NodeTree represents a node within the binary tree.
type NodeTree struct {
	Data  int
	Left  *NodeTree
	Right *NodeTree
}

// Tree represents the main structure of a binary tree.
type Tree struct {
	Raiz *NodeTree
}

// NewArbol creates and returns a new empty tree.
// Complexity: O(1)
func NewArbol() *Tree {
	return &Tree{Raiz: nil}
}

// Search looks for a value in the tree and returns true if found.
// Complexity: O(n)
func (tree *Tree) Search(node *NodeTree, valor int) bool {
	if node == nil {
		return false
	}
	if node.Data == valor {
		return true
	} else if valor < node.Data {
		return tree.Search(node.Left, valor)
	} else {
		return tree.Search(node.Right, valor)
	}
}

// InsertNode inserts a new node into the binary search tree.
// Complexity: O(n)
func (tree *Tree) InsertNode(node **NodeTree, valor int) {
	if *node == nil {
		*node = &NodeTree{Data: valor}
		return
	}
	if valor < (*node).Data {
		tree.InsertNode(&((*node).Left), valor)
	} else if valor > (*node).Data {
		tree.InsertNode(&((*node).Right), valor)
	}
}

// DeleteNode removes a specific node from the tree.
// Complexity: O(n)
func (tree *Tree) DeleteNode(node *NodeTree, valor int) *NodeTree {
	if node == nil {
		return nil
	}

	if valor < node.Data {
		node.Left = tree.DeleteNode(node.Left, valor)
	} else if valor > node.Data {
		node.Right = tree.DeleteNode(node.Right, valor)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minimo := tree.FindMin(node.Right)
		node.Data = minimo.Data
		node.Right = tree.DeleteNode(node.Right, minimo.Data)
	}
	return node
}

// DeleteSubTree removes an entire subtree starting from the given node.
// Complexity: O(n)
func (tree *Tree) DeleteSubTree(node *NodeTree) {
	if node == nil {
		return
	}
	tree.DeleteSubTree(node.Left)
	tree.DeleteSubTree(node.Right)
	node.Left = nil
	node.Right = nil
}

// FindMin searches and returns the node with the minimum value in the tree.
// Complexity: O(h), worst case O(n)
func (tree *Tree) FindMin(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return tree.FindMin(node.Left)
}

// FindMax searches and returns the node with the maximum value in the tree.
// Complexity: O(h), worst case O(n)
func (tree *Tree) FindMax(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node
	}
	return tree.FindMax(node.Right)
}

// DeleteMin removes the node with the minimum value from the tree.
// Complexity: O(h), worst case O(n)
func (tree *Tree) DeleteMin(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node.Right
	}
	node.Left = tree.DeleteMin(node.Left)
	return node
}

// DeleteMax removes the node with the maximum value from the tree.
// Complexity: O(h), worst case O(n)
func (tree *Tree) DeleteMax(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node.Left
	}
	node.Right = tree.DeleteMax(node.Right)
	return node
}

// Height calculates and returns the height of the tree.
// Complexity: O(n)
func (tree *Tree) Height(node *NodeTree) int {
	if node == nil {
		return 0
	}
	left := tree.Height(node.Left)
	right := tree.Height(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// CountNodes returns the total number of nodes in the tree.
// Complexity: O(n)
func (tree *Tree) CountNodes(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return 1 + tree.CountNodes(node.Left) + tree.CountNodes(node.Right)
}

// Weight returns the sum of all values stored in the tree.
// Complexity: O(n)
func (tree *Tree) Weight(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return node.Data + tree.Weight(node.Left) + tree.Weight(node.Right)
}

// PrintInorder prints the tree in inorder traversal (left, root, right).
// Complexity: O(n)
func (tree *Tree) PrintInorder(node *NodeTree) {
	if node != nil {
		tree.PrintInorder(node.Left)
		fmt.Printf("%d ", node.Data)
		tree.PrintInorder(node.Right)
	}
}

// PrintPostorder prints the tree in postorder traversal (left, right, root).
// Complexity: O(n)
func (tree *Tree) PrintPostorder(node *NodeTree) {
	if node != nil {
		tree.PrintPostorder(node.Left)
		tree.PrintPostorder(node.Right)
		fmt.Printf("%d ", node.Data)
	}
}

// PrintPreorder prints the tree in preorder traversal (root, left, right).
// Complexity: O(n)
func (tree *Tree) PrintPreorder(node *NodeTree) {
	if node != nil {
		fmt.Printf("%d ", node.Data)
		tree.PrintPreorder(node.Left)
		tree.PrintPreorder(node.Right)
	}
}
