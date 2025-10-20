package lists

import "fmt"

// NodeTree represents a single node in the binary tree.
type NodeTree struct {
	Data  int
	Left  *NodeTree
	Right *NodeTree
}

// Tree represents a binary search tree.
type Tree struct {
	Root *NodeTree
}

// NewTree creates and returns an empty binary search tree.
func NewTree() *Tree {
	return &Tree{Root: nil}
}

// Insert inserts a new value into the binary search tree.
func (t *Tree) Insert(value int) {
	t.Root = insertRecursive(t.Root, value)
}

func insertRecursive(node *NodeTree, value int) *NodeTree {
	if node == nil {
		return &NodeTree{Data: value}
	}
	if value < node.Data {
		node.Left = insertRecursive(node.Left, value)
	} else if value > node.Data {
		node.Right = insertRecursive(node.Right, value)
	}
	return node
}

// Search returns true if the given value exists in the tree.
func (t *Tree) Search(value int) bool {
	return searchRecursive(t.Root, value)
}

func searchRecursive(node *NodeTree, value int) bool {
	if node == nil {
		return false
	}
	if node.Data == value {
		return true
	} else if value < node.Data {
		return searchRecursive(node.Left, value)
	}
	return searchRecursive(node.Right, value)
}

// Delete removes a node with the specified value from the tree.
func (t *Tree) Delete(value int) {
	t.Root = deleteRecursive(t.Root, value)
}

func deleteRecursive(node *NodeTree, value int) *NodeTree {
	if node == nil {
		return nil
	}

	if value < node.Data {
		node.Left = deleteRecursive(node.Left, value)
	} else if value > node.Data {
		node.Right = deleteRecursive(node.Right, value)
	} else {
		// Node found
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		// Replace with the minimum node from the right subtree
		minNode := findMin(node.Right)
		node.Data = minNode.Data
		node.Right = deleteRecursive(node.Right, minNode.Data)
	}
	return node
}

// FindMin returns the node with the minimum value in the tree.
func (t *Tree) FindMin() *NodeTree {
	return findMin(t.Root)
}

func findMin(node *NodeTree) *NodeTree {
	if node == nil || node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

// FindMax returns the node with the maximum value in the tree.
func (t *Tree) FindMax() *NodeTree {
	return findMax(t.Root)
}

func findMax(node *NodeTree) *NodeTree {
	if node == nil || node.Right == nil {
		return node
	}
	return findMax(node.Right)
}

// Height returns the height of the tree.
func (t *Tree) Height() int {
	return heightRecursive(t.Root)
}

func heightRecursive(node *NodeTree) int {
	if node == nil {
		return 0
	}
	left := heightRecursive(node.Left)
	right := heightRecursive(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// CountNodes returns the total number of nodes in the tree.
func (t *Tree) CountNodes() int {
	return countNodesRecursive(t.Root)
}

func countNodesRecursive(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return 1 + countNodesRecursive(node.Left) + countNodesRecursive(node.Right)
}

// Sum returns the sum of all values in the tree.
func (t *Tree) Sum() int {
	return sumRecursive(t.Root)
}

func sumRecursive(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return node.Data + sumRecursive(node.Left) + sumRecursive(node.Right)
}

// Clear removes all nodes from the tree.
func (t *Tree) Clear() {
	t.Root = nil
}

// PrintInorder prints the tree in inorder traversal (left, root, right).
func (t *Tree) PrintInorder() {
	printInorderRecursive(t.Root)
	fmt.Println()
}

func printInorderRecursive(node *NodeTree) {
	if node != nil {
		printInorderRecursive(node.Left)
		fmt.Printf("%d ", node.Data)
		printInorderRecursive(node.Right)
	}
}

// PrintPreorder prints the tree in preorder traversal (root, left, right).
func (t *Tree) PrintPreorder() {
	printPreorderRecursive(t.Root)
	fmt.Println()
}

func printPreorderRecursive(node *NodeTree) {
	if node != nil {
		fmt.Printf("%d ", node.Data)
		printPreorderRecursive(node.Left)
		printPreorderRecursive(node.Right)
	}
}

// PrintPostorder prints the tree in postorder traversal (left, right, root).
func (t *Tree) PrintPostorder() {
	printPostorderRecursive(t.Root)
	fmt.Println()
}

func printPostorderRecursive(node *NodeTree) {
	if node != nil {
		printPostorderRecursive(node.Left)
		printPostorderRecursive(node.Right)
		fmt.Printf("%d ", node.Data)
	}
}
