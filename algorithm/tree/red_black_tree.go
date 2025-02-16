package tree

const (
	RED   = true
	BLACK = false
)

// Узел красно-черного дерева
type NodeRBT struct {
	Value  int
	Color  bool // true - красный, false - черный
	Left   *NodeRBT
	Right  *NodeRBT
	Parent *NodeRBT
}

// Красно-черное дерево
type RedBlackTree struct {
	Root *NodeRBT
}

// Создание нового узла
func NewNode(value int, color bool) *NodeRBT {
	return &NodeRBT{
		Value: value,
		Color: color,
	}
}

// Вставка значения в дерево
func (t *RedBlackTree) Insert(value int) {
	newNode := NewNode(value, RED)
	if t.Root == nil {
		t.Root = newNode
	} else {
		t.insertNode(t.Root, newNode)
	}
	t.fixInsert(newNode)
}

// Вспомогательная функция для вставки узла
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

// Балансировка дерева после вставки
func (t *RedBlackTree) fixInsert(node *NodeRBT) {
	for node.Parent != nil && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle != nil && uncle.Color == RED {
				// красный
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					// черный, узел — правый потомок
					node = node.Parent
					t.rotateLeft(node)
				}
				// черный, узел — левый потомок
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.rotateRight(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle != nil && uncle.Color == RED {
				// красный
				node.Parent.Color = BLACK
				uncle.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					// черный, узел — левый потомок
					node = node.Parent
					t.rotateRight(node)
				}
				// черный, узел — правый потомок
				node.Parent.Color = BLACK
				node.Parent.Parent.Color = RED
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

// Левый поворот
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

// Правый поворот
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

// Поиск значения в дереве
func (t *RedBlackTree) Search(value int) bool {
	return t.searchNode(t.Root, value)
}

// Вспомогательная функция для поиска
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

// Обход дерева in-order и возврат результата в виде слайса
func (t *RedBlackTree) InOrderTraversal() []int {
	result := []int{}
	t.inOrder(t.Root, &result)
	return result
}

// Рекурсивная функция для обхода in-order
func (t *RedBlackTree) inOrder(node *NodeRBT, result *[]int) {
	if node != nil {
		t.inOrder(node.Left, result)          // Рекурсивно обходим левое поддерево
		*result = append(*result, node.Value) // Добавляем значение текущего узла
		t.inOrder(node.Right, result)         // Рекурсивно обходим правое поддерево
	}
}
