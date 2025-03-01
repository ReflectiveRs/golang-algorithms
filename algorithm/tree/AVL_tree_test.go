package tree

import (
	"reflect"
	"testing"
)

func TestAVLTree(t *testing.T) {
	tree := &AVLTree{}
	values := []int{30, 20, 40, 10, 25, 35, 50}

	// Inserting values
	for _, v := range values {
		tree.Insert(v)
	}

	// Checking the search
	tests := []struct {
		value  int
		exists bool
	}{
		{30, true},
		{20, true},
		{40, true},
		{10, true},
		{25, true},
		{35, true},
		{50, true},
		{5, false},
	}

	for _, test := range tests {
		result := tree.Search(test.value)
		if result != test.exists {
			t.Errorf("Search(%d) = %v; expected %v", test.value, result, test.exists)
		}
	}

	// Checking symmetric traversal
	expectedInOrder := []int{10, 20, 25, 30, 35, 40, 50}
	resultInOrder := tree.InOrderTraversal()
	if !reflect.DeepEqual(resultInOrder, expectedInOrder) {
		t.Errorf("InOrderTraversal() = %v; expected %v", resultInOrder, expectedInOrder)
	}
}

func TestAVLTreeInsertionAndBalance(t *testing.T) {
	tree := &AVLTree{}

	// Inserting values that should trigger balancing
	values := []int{10, 20, 30, 40, 50, 25}
	for _, v := range values {
		tree.Insert(v)
	}

	// Checking symmetric traversal
	expectedInOrder := []int{10, 20, 25, 30, 40, 50}
	resultInOrder := tree.InOrderTraversal()
	if !reflect.DeepEqual(resultInOrder, expectedInOrder) {
		t.Errorf("InOrderTraversal() = %v; expected %v", resultInOrder, expectedInOrder)
	}
}

func TestAVLTreeSearchNonExistent(t *testing.T) {
	tree := &AVLTree{}
	values := []int{50, 30, 70}
	for _, v := range values {
		tree.Insert(v)
	}

	// Checking the search for non-existent value
	if tree.Search(100) {
		t.Error("Expected Search(100) to be false, got true")
	}
}
