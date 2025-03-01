package tree

const (
	RED   = true
	BLACK = false
)

// Node red-black tree
type NodeRBT struct {
	Value  int
	Color  bool // true - red, false - black
	Left   *NodeRBT
	Right  *NodeRBT
	Parent *NodeRBT
}

type RedBlackTree struct {
	Root *NodeRBT
}

func NewNode(value int, color bool) *NodeRBT {
	return &NodeRBT{
		Value: value,
		Color: color,
	}
}

func (t *RedBlackTree) Insert(value int) {
	newNode := NewNode(value, RED)
	if t.Root == nil {
		t.Root = newNode
	} else {
		t.insertNode(t.Root, newNode)
	}
	t.fixInsert(newNode)
}

// Auxiliary function for inserting a node
func (t *RedBlackTree) insertNode(root, newNode *NodeRBT) {
	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
			newNode.Parent = root
		} else {
			t.insertNode(root.Left, newNode)
		}
	} else {
		if root.Right == nil {
			root.Right = newNode
			newNode.Parent = root
		} else {
			t.insertNode(root.Right, newNode)
		}
	}
}

// Balancing tree after insertion
func (t *RedBlackTree) fixInsert(node *NodeRBT) {
	for node.Parent != nil && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RED {
				// Red
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					// Black, node is the right descendant
					node = node.Parent
					t.rotateLeft(node)
				}
				// Black, node is the left descendant
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RED {
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					// Black, node is the left descendant
					node = node.Parent
					t.rotateRight(node)
				}
				// Black, node is the right descendant
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

// Left turn
func (t *RedBlackTree) rotateLeft(node *NodeRBT) {
	rightChild := node.Right
	node.Right = rightChild.Left
	if rightChild.Left != nil {
		rightChild.Left.Parent = node
	}
	rightChild.Parent = node.Parent
	if node.Parent == nil {
		t.Root = rightChild
	} else if node == node.Parent.Left {
		node.Parent.Left = rightChild
	} else {
		node.Parent.Right = rightChild
	}
	rightChild.Left = node
	node.Parent = rightChild
}

// Right turn
func (t *RedBlackTree) rotateRight(node *NodeRBT) {
	leftChild := node.Left
	node.Left = leftChild.Right
	if leftChild.Right != nil {
		leftChild.Right.Parent = node
	}
	leftChild.Parent = node.Parent
	if node.Parent == nil {
		t.Root = leftChild
	} else if node == node.Parent.Right {
		node.Parent.Right = leftChild
	} else {
		node.Parent.Left = leftChild
	}
	leftChild.Right = node
	node.Parent = leftChild
}

// Searching for a value in a tree
func (t *RedBlackTree) Search(value int) bool {
	return t.searchNode(t.Root, value)
}

// Searching for a node in a tree
func (t *RedBlackTree) searchNode(node *NodeRBT, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return t.searchNode(node.Left, value)
	} else if value > node.Value {
		return t.searchNode(node.Right, value)
	}
	return true
}

// Traversing the in-order tree and returning the result as a slice
func (t *RedBlackTree) InOrderTraversal() []int {
	result := []int{}
	t.inOrder(t.Root, &result)
	return result
}

// Recursive function for in-order traversal
func (t *RedBlackTree) inOrder(node *NodeRBT, result *[]int) {
	if node != nil {
		t.inOrder(node.Left, result)          // Recursively traversing the left subtree
		*result = append(*result, node.Value) // Adding the value of the current node
		t.inOrder(node.Right, result)         // Recursively traversing the right subtree
	}
}
