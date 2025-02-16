package tree

// Node представляет узел AVL-дерева
type AVLNode struct {
	Value  int
	Left   *AVLNode
	Right  *AVLNode
	Height int
}

// AVLTree представляет AVL-дерево
type AVLTree struct {
	Root *AVLNode
}

// GetHeight возвращает высоту узла
func (n *AVLNode) GetHeight() int {
	if n == nil {
		return 0
	}
	return n.Height
}

// UpdateHeight обновляет высоту узла
func (n *AVLNode) UpdateHeight() {
	n.Height = max(n.Left.GetHeight(), n.Right.GetHeight()) + 1
}

// max возвращает максимальное из двух значений
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// GetBalance возвращает баланс узла
func (n *AVLNode) GetBalance() int {
	if n == nil {
		return 0
	}
	return n.Left.GetHeight() - n.Right.GetHeight()
}

// RightRotate выполняет правый поворот вокруг узла
func RightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.UpdateHeight()
	x.UpdateHeight()

	return x
}

// LeftRotate выполняет левый поворот вокруг узла
func LeftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.UpdateHeight()
	y.UpdateHeight()

	return y
}

// Insert вставляет новое значение в AVL-дерево
func (tree *AVLTree) Insert(value int) {
	tree.Root = tree.Root.insert(value)
}

// insert рекурсивно вставляет новое значение в поддерево
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

	// Левый левый случай
	if balance > 1 && value < n.Left.Value {
		return RightRotate(n)
	}

	// Правый правый случай
	if balance < -1 && value > n.Right.Value {
		return LeftRotate(n)
	}

	// Левый правый случай
	if balance > 1 && value > n.Left.Value {
		n.Left = LeftRotate(n.Left)
		return RightRotate(n)
	}

	// Правый левый случай
	if balance < -1 && value < n.Right.Value {
		n.Right = RightRotate(n.Right)
		return LeftRotate(n)
	}

	return n
}

// Search ищет значение в AVL-дереве
func (tree *AVLTree) Search(value int) bool {
	return tree.Root.search(value)
}

// search рекурсивно ищет значение в поддереве
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

// InOrderTraversal выполняет обход дерева в симметричном порядке
func (tree *AVLTree) InOrderTraversal() []int {
	var result []int
	tree.Root.inOrderTraversal(&result)
	return result
}

// inOrderTraversal рекурсивно выполняет симметричный обход поддерева
func (n *AVLNode) inOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.inOrderTraversal(result)
		*result = append(*result, n.Value)
		n.Right.inOrderTraversal(result)
	}
}
