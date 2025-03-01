package tree

import "testing"

// Checking that the tree satisfies the properties of a red-black tree
func (t *RedBlackTree) checkRBProperties(node *NodeRBT) bool {
	if node == nil {
		return true
	}

	// root is always black
	if node == t.Root && node.Color != BLACK {
		return false
	}

	// If node is red, then both of its descendants are black.
	if node.Color == RED {
		if (node.Left != nil && node.Left.Color != BLACK) || (node.Right != nil && node.Right.Color != BLACK) {
			return false
		}
	}

	// Recursive check for left and right subtrees
	return t.checkRBProperties(node.Left) && t.checkRBProperties(node.Right)
}

// Calculating the black height of a tree
func (t *RedBlackTree) blackHeight(node *NodeRBT) int {
	if node == nil {
		return 1
	}

	leftBlackHeight := t.blackHeight(node.Left)
	rightBlackHeight := t.blackHeight(node.Right)

	// All paths from the node to the leaves contain the same number of black nodes.
	if leftBlackHeight != rightBlackHeight {
		return -1 // Invalid tree
	}

	if node.Color == BLACK {
		return leftBlackHeight + 1
	}
	return leftBlackHeight
}

// Check for insertion and checking properties of red-black tree
func TestInsertAndRBProperties(t *testing.T) {
	tree := &RedBlackTree{}

	// Inserting elements
	values := []int{10, 20, 30, 15, 25, 5}
	for _, value := range values {
		tree.Insert(value)
	}

	// Checking the properties of the red-black tree
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Дерево не удовлетворяет свойствам красно-черного дерева после вставки")
	}

	// Checking the black height
	if tree.blackHeight(tree.Root) == -1 {
		t.Error("Черная высота дерева не совпадает для всех путей")
	}
}

// Checking the search for elements
func TestSearch(t *testing.T) {
	tree := &RedBlackTree{}

	// Inserting elements
	values := []int{10, 20, 30, 15, 25, 5}
	for _, value := range values {
		tree.Insert(value)
	}

	// Checking for inserted elements
	for _, value := range values {
		if !tree.Search(value) {
			t.Errorf("Ожидалось, что значение %d будет найдено в дереве", value)
		}
	}

	// Check for the absence of elements that are not in the tree
	nonExistentValues := []int{1, 100, 42}
	for _, value := range nonExistentValues {
		if tree.Search(value) {
			t.Errorf("Ожидалось, что значение %d не будет найдено в дереве", value)
		}
	}
}

// Checking if the tree is empty
func TestEmptyTree(t *testing.T) {
	tree := &RedBlackTree{}

	// Check that searching in an empty tree returns false
	if tree.Search(10) {
		t.Error("Ожидалось, что поиск в пустом дереве вернет false")
	}

	// Checking the red-black tree properties for an empty tree
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Пустое дерево должно удовлетворять свойствам красно-черного дерева")
	}
}

// Check for inserting duplicates
func TestInsertDuplicates(t *testing.T) {
	tree := &RedBlackTree{}

	// Insert duplicates
	values := []int{10, 10, 20, 20, 30, 30}
	for _, value := range values {
		tree.Insert(value)
	}

	// Check that the tree handled duplicates correctly
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Дерево не удовлетворяет свойствам красно-черного дерева после вставки дубликатов")
	}

	// Checking the black height
	if tree.blackHeight(tree.Root) == -1 {
		t.Error("Черная высота дерева не совпадает для всех путей после вставки дубликатов")
	}
}
