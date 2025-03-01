package tree

// Node of a binary tree
type NodeBST struct {
	Value int
	Left  *NodeBST
	Right *NodeBST
}

// Binary search tree
type BST struct {
	Root *NodeBST
}

// Adds a value to the binary search tree
func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &NodeBST{Value: value}
	} else {
		bst.Root.insert(value)
	}
}

// Recursively inserts a value into a subtree
func (n *NodeBST) insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &NodeBST{Value: value}
		} else {
			n.Left.insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &NodeBST{Value: value}
		} else {
			n.Right.insert(value)
		}
	}
}

// Searches for a value in a binary search tree
func (bst *BST) Search(value int) bool {
	return bst.Root.search(value)
}

// Recursively searches for a value in a subtree
func (n *NodeBST) search(value int) bool {
	if n == nil {
		return false
	}
	if value == n.Value {
		return true
	} else if value < n.Value {
		return n.Left.search(value)
	} else {
		return n.Right.search(value)
	}
}

// Performs a tree traversal in a symmetrical order
func (bst *BST) InOrderTraversal() []int {
	var result []int
	bst.Root.inOrderTraversal(&result)
	return result
}

// Recursively performs symmetric traversal of a subtree
func (n *NodeBST) inOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.inOrderTraversal(result)
		*result = append(*result, n.Value)
		n.Right.inOrderTraversal(result)
	}
}

// Performs a preliminary tree traversal
func (bst *BST) PreOrderTraversal() []int {
	var result []int
	bst.Root.preOrderTraversal(&result)
	return result
}

// Recursively performs a preliminary traversal of the subtree
func (n *NodeBST) preOrderTraversal(result *[]int) {
	if n != nil {
		*result = append(*result, n.Value) // First we add the root
		n.Left.preOrderTraversal(result)   // Then the left subtree
		n.Right.preOrderTraversal(result)  // And the right subtree
	}
}

// Performs a tree traversal in the postmodern order
func (bst *BST) PostOrderTraversal() []int {
	var result []int
	bst.Root.postOrderTraversal(&result)
	return result
}

// Recursively performs a postmodern subtree traversal
func (n *NodeBST) postOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.postOrderTraversal(result)  // First, the left subtree
		n.Right.postOrderTraversal(result) // Then the right subtree
		*result = append(*result, n.Value) // And add the root at the end.
	}
}
