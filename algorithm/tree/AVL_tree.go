package tree

// Node AVL tree
type AVLNode struct {
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// AVLTree represents an AVL tree
type AVLTree struct {
	Root *AVLNode
}

// Get Node height
func (n *AVLNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Update Node heigh
func (n *AVLNode) UpdateHeight() {
	n.Height = max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
}

// Returns the maximum of two values
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Returns node balance
func (n *AVLNode) GetBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

// Performs right turn around the node
func RightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.UpdateHeight()
	x.UpdateHeight()

	return x
}

// Performs a left turn around the node
func LeftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.UpdateHeight()
	y.UpdateHeight()

	return y
}

// Inserts new value into the AVL tree
func (tree *AVLTree) Insert(value int) {
	tree.Root = tree.Root.insert(value)
}

// Recursively inserts new value into a subtree
func (n *AVLNode) insert(value int) *AVLNode {
	if n == nil {
		return &AVLNode{Value: value, Height: 1}
	}

	if value < n.Value {
		n.Left = n.Left.insert(value)
	} else {
		n.Right = n.Right.insert(value)
	}

	n.UpdateHeight()

	balance := n.GetBalance()

	// Left left case
	if balance > 1 && value < n.Left.Value {
		return RightRotate(n)
	}

	// Right right case
	if balance < -1 && value > n.Right.Value {
		return LeftRotate(n)
	}

	// Left-right case
	if balance > 1 && value > n.Left.Value {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}

	// Right-left case
	if balance < -1 && value < n.Right.Value {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}

	return n
}

// Searches for value in the AVL tree
func (tree *AVLTree) Search(value int) bool {
	return tree.Root.search(value)
}

// Recursively searches for value in subtree
func (n *AVLNode) search(value int) bool {
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

// Performs tree traversal in symmetrical order
func (tree *AVLTree) InOrderTraversal() []int {
	var result []int
	tree.Root.inOrderTraversal(&result)
	return result
}

// Recursively performs symmetric traversal of subtree
func (n *AVLNode) inOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.inOrderTraversal(result)
		*result = append(*result, n.Value)
		n.Right.inOrderTraversal(result)
	}
}
