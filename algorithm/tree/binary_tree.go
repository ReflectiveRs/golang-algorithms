package tree

// представляет узел бинарного дерева
type NodeBST struct {
	Value int
	Left  *NodeBST
	Right *NodeBST
}

// представляет бинарное дерево поиска
type BST struct {
	Root *NodeBST
}

// добавляет значение в бинарное дерево поиска
func (bst *BST) Insert(value int) {
	if bst.Root == nil {
		bst.Root = &NodeBST{Value: value}
	} else {
		bst.Root.insert(value)
	}
}

// рекурсивно вставляет значение в поддерево
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

// ищет значение в бинарном дереве поиска
func (bst *BST) Search(value int) bool {
	return bst.Root.search(value)
}

// рекурсивно ищет значение в поддереве
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

// выполняет обход дерева в симметричном порядке
func (bst *BST) InOrderTraversal() []int {
	var result []int
	bst.Root.inOrderTraversal(&result)
	return result
}

// рекурсивно выполняет симметричный обход поддерева
func (n *NodeBST) inOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.inOrderTraversal(result)
		*result = append(*result, n.Value)
		n.Right.inOrderTraversal(result)
	}
}

// выполняет обход дерева в предварительном порядке
func (bst *BST) PreOrderTraversal() []int {
	var result []int
	bst.Root.preOrderTraversal(&result)
	return result
}

// рекурсивно выполняет предварительный обход поддерева
func (n *NodeBST) preOrderTraversal(result *[]int) {
	if n != nil {
		*result = append(*result, n.Value) // Сначала добавляем корень
		n.Left.preOrderTraversal(result)   // Затем левое поддерево
		n.Right.preOrderTraversal(result)  // И правое поддерево
	}
}

// выполняет обход дерева в постордерном порядке
func (bst *BST) PostOrderTraversal() []int {
	var result []int
	bst.Root.postOrderTraversal(&result)
	return result
}

// рекурсивно выполняет постордерный обход поддерева
func (n *NodeBST) postOrderTraversal(result *[]int) {
	if n != nil {
		n.Left.postOrderTraversal(result)  // Сначала левое поддерево
		n.Right.postOrderTraversal(result) // Затем правое поддерево
		*result = append(*result, n.Value) // И добавляем корень в конце
	}
}
