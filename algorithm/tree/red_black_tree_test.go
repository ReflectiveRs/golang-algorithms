package tree

import "testing"

// Проверка, что дерево удовлетворяет свойствам красно-черного дерева
func (t *RedBlackTree) checkRBProperties(node *NodeRBT) bool {
	if node == nil {
		return true
	}

	// Корень всегда черный
	if node == t.Root && node.Color != BLACK {
		return false
	}

	// Если узел красный, то оба его потомка черные
	if node.Color == RED {
		if (node.Left != nil && node.Left.Color != BLACK) || (node.Right != nil && node.Right.Color != BLACK) {
			return false
		}
	}

	// Рекурсивная проверка для левого и правого поддеревьев
	return t.checkRBProperties(node.Left) && t.checkRBProperties(node.Right)
}

// Подсчет черной высоты дерева
func (t *RedBlackTree) blackHeight(node *NodeRBT) int {
	if node == nil {
		return 1
	}

	leftBlackHeight := t.blackHeight(node.Left)
	rightBlackHeight := t.blackHeight(node.Right)

	// Все пути от узла до листьев содержат одинаковое количество черных узлов
	if leftBlackHeight != rightBlackHeight {
		return -1 // Невалидное дерево
	}

	if node.Color == BLACK {
		return leftBlackHeight + 1
	}
	return leftBlackHeight
}

// Тест для вставки и проверки свойств красно-черного дерева
func TestInsertAndRBProperties(t *testing.T) {
	tree := &RedBlackTree{}

	// Вставляем элементы
	values := []int{10, 20, 30, 15, 25, 5}
	for _, value := range values {
		tree.Insert(value)
	}

	// Проверяем свойства красно-черного дерева
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Дерево не удовлетворяет свойствам красно-черного дерева после вставки")
	}

	// Проверяем черную высоту
	if tree.blackHeight(tree.Root) == -1 {
		t.Error("Черная высота дерева не совпадает для всех путей")
	}
}

// Тест для поиска элементов
func TestSearch(t *testing.T) {
	tree := &RedBlackTree{}

	// Вставляем элементы
	values := []int{10, 20, 30, 15, 25, 5}
	for _, value := range values {
		tree.Insert(value)
	}

	// Проверяем наличие вставленных элементов
	for _, value := range values {
		if !tree.Search(value) {
			t.Errorf("Ожидалось, что значение %d будет найдено в дереве", value)
		}
	}

	// Проверяем отсутствие элементов, которых нет в дереве
	nonExistentValues := []int{1, 100, 42}
	for _, value := range nonExistentValues {
		if tree.Search(value) {
			t.Errorf("Ожидалось, что значение %d не будет найдено в дереве", value)
		}
	}
}

// Тест для пустого дерева
func TestEmptyTree(t *testing.T) {
	tree := &RedBlackTree{}

	// Проверяем, что поиск в пустом дереве возвращает false
	if tree.Search(10) {
		t.Error("Ожидалось, что поиск в пустом дереве вернет false")
	}

	// Проверяем свойства красно-черного дерева для пустого дерева
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Пустое дерево должно удовлетворять свойствам красно-черного дерева")
	}
}

// Тест для вставки дубликатов
func TestInsertDuplicates(t *testing.T) {
	tree := &RedBlackTree{}

	// Вставляем дубликаты
	values := []int{10, 10, 20, 20, 30, 30}
	for _, value := range values {
		tree.Insert(value)
	}

	// Проверяем, что дерево корректно обработало дубликаты
	if !tree.checkRBProperties(tree.Root) {
		t.Error("Дерево не удовлетворяет свойствам красно-черного дерева после вставки дубликатов")
	}

	// Проверяем черную высоту
	if tree.blackHeight(tree.Root) == -1 {
		t.Error("Черная высота дерева не совпадает для всех путей после вставки дубликатов")
	}
}
