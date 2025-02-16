package tree

import (
	"reflect"
	"testing"
)

func TestBST(t *testing.T) {
	bst := &BST{}
	values := []int{10, 5, 15, 3, 7, 12, 18}

	for _, v := range values {
		bst.Insert(v)
	}

	tests := []struct {
		value  int
		exists bool
	}{
		{10, true},
		{5, true},
		{15, true},
		{3, true},
		{7, true},
		{12, true},
		{18, true},
		{20, false},
	}

	for _, test := range tests {
		result := bst.Search(test.value)
		if result != test.exists {
			t.Errorf("Search(%d) = %v; expected %v", test.value, result, test.exists)
		}
	}

	// Проверка симметричного обхода
	expectedTraversal := []int{3, 5, 7, 10, 12, 15, 18}
	resultTraversal := bst.InOrderTraversal()
	if !reflect.DeepEqual(resultTraversal, expectedTraversal) {
		t.Errorf("InOrderTraversal() = %v; expected %v", resultTraversal, expectedTraversal)
	}

	// Проверка предварительного обхода
	expectedPreOrder := []int{10, 5, 3, 7, 15, 12, 18}
	resultPreOrder := bst.PreOrderTraversal()
	if !reflect.DeepEqual(resultPreOrder, expectedPreOrder) {
		t.Errorf("PreOrderTraversal() = %v; expected %v", resultPreOrder, expectedPreOrder)
	}

	// Проверка постордерном обхода
	expectedPostOrder := []int{10, 5, 3, 7, 15, 12, 18}
	resultPostOrder := bst.PreOrderTraversal()
	if !reflect.DeepEqual(resultPostOrder, expectedPostOrder) {
		t.Errorf("PreOrderTraversal() = %v; expected %v", resultPostOrder, expectedPostOrder)
	}
}
