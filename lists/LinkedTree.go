package lists

import "fmt"

// NodeTree represents a single node in the binary search tree.
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
// Complexity: O(1)
func NewTree() *Tree {
	return &Tree{Root: nil}
}

// Search searches for a value in the tree and returns true if found.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) Search(value int) bool {
	return searchRecursive(t.Root, value)
}

func searchRecursive(node *NodeTree, value int) bool {
	if node == nil {
		return false
	}
	if value == node.Data {
		return true
	} else if value < node.Data {
		return searchRecursive(node.Left, value)
	}
	return searchRecursive(node.Right, value)
}

// InsertNode inserts a new node into the binary search tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) InsertNode(value int) {
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

// DeleteNode removes a node with the specified value from the tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) DeleteNode(value int) {
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
		minNode := findMin(node.Right)
		node.Data = minNode.Data
		node.Right = deleteRecursive(node.Right, minNode.Data)
	}
	return node
}

// DeleteSubTree removes the entire subtree starting from the root.
// Complexity: O(n)
func (t *Tree) DeleteSubTree() {
	deleteSubTreeRecursive(t.Root)
	t.Root = nil
}

func deleteSubTreeRecursive(node *NodeTree) {
	if node == nil {
		return
	}
	deleteSubTreeRecursive(node.Left)
	deleteSubTreeRecursive(node.Right)
	node.Left = nil
	node.Right = nil
}

// FindMin returns the node with the smallest value in the tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) FindMin() *NodeTree {
	return findMin(t.Root)
}

func findMin(node *NodeTree) *NodeTree {
	if node == nil || node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

// FindMax returns the node with the largest value in the tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) FindMax() *NodeTree {
	return findMax(t.Root)
}

func findMax(node *NodeTree) *NodeTree {
	if node == nil || node.Right == nil {
		return node
	}
	return findMax(node.Right)
}

// DeleteMin deletes the node with the smallest value from the tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) DeleteMin() {
	t.Root = deleteMinRecursive(t.Root)
}

func deleteMinRecursive(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMinRecursive(node.Left)
	return node
}

// DeleteMax deletes the node with the largest value from the tree.
// Complexity: O(log n) average, O(n) worst.
func (t *Tree) DeleteMax() {
	t.Root = deleteMaxRecursive(t.Root)
}

func deleteMaxRecursive(node *NodeTree) *NodeTree {
	if node == nil {
		return nil
	}
	if node.Right == nil {
		return node.Left
	}
	node.Right = deleteMaxRecursive(node.Right)
	return node
}

// Height returns the height of the tree.
// Complexity: O(n)
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
// Complexity: O(n)
func (t *Tree) CountNodes() int {
	return countNodesRecursive(t.Root)
}

func countNodesRecursive(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return 1 + countNodesRecursive(node.Left) + countNodesRecursive(node.Right)
}

// Weight returns the sum of all node values in the tree.
// Complexity: O(n)
func (t *Tree) Weight() int {
	return weightRecursive(t.Root)
}

func weightRecursive(node *NodeTree) int {
	if node == nil {
		return 0
	}
	return node.Data + weightRecursive(node.Left) + weightRecursive(node.Right)
}

// PrintInorder prints the tree contents in inorder traversal (left, root, right).
// Complexity: O(n)
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

// PrintPostorder prints the tree contents in postorder traversal (left, right, root).
// Complexity: O(n)
func (t *Tree) PrintPostorder() {
	printPostorderRecursive(t.Root)
	fmt.Println()
}

func printPostorderRecursive(node *NodeTree) {
	if node != nil {
		printPostorderRecursive(node.Left)
		printPostorderRecursive(node.Right)
		fmt.Printf("%d\n", node.Data)
	}
}

// PrintPreorder prints the tree contents in preorder traversal (root, left, right).
// Complexity: O(n)
func (t *Tree) PrintPreorder() {
	printPreorderRecursive(t.Root)
	fmt.Println()
}

func printPreorderRecursive(node *NodeTree) {
	if node != nil {
		fmt.Printf("%d\n ", node.Data)
		printPreorderRecursive(node.Left)
		printPreorderRecursive(node.Right)
	}
}
