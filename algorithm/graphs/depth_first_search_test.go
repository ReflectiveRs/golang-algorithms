package graphs

import "testing"

func TestDFSRecursive(t *testing.T) {
	graph := NewGraphDFS()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	tests := []struct {
		start    int
		expected []int
	}{
		{0, []int{0, 1, 3, 4, 2, 5, 6}}, // Expected order
		{1, []int{1, 3, 4}},             // Expected order
		{2, []int{2, 5, 6}},             // Expected order
	}

	for _, test := range tests {
		visited := make(map[int]bool)
		var result []int
		graph.DFSRecursive(test.start, visited, &result)

		// We check that the length of the result matches the expected one
		if len(result) != len(test.expected) {
			t.Errorf("DFSRecursive(%d) = %v; expected length %d", test.start, result, len(test.expected))
		}

		// We check that the result contains all expected elements
		if !containsAll(result, test.expected) {
			t.Errorf("DFSRecursive(%d) = %v; expected %v", test.start, result, test.expected)
		}
	}
}

// containsAll checks whether one slice contains all elements of another
func containsAll(a, b []int) bool {
	m := make(map[int]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if !m[item] {
			return false
		}
	}
	return true
}

func TestDFSIterative(t *testing.T) {
	graph := NewGraphDFS()
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)

	tests := []struct {
		start    int
		expected []int
	}{
		{0, []int{0, 1, 3, 4, 2, 6, 5}}, // Expected order
		{1, []int{1, 3, 4}},             // Expected order
		{2, []int{2, 5, 6}},             // Expected order
	}

	for _, test := range tests {
		result := graph.DFSIterative(test.start)

		// We check that the length of the result matches the expected one.
		if len(result) != len(test.expected) {
			t.Errorf("DFSIterative(%d) = %v; expected length %d", test.start, result, len(test.expected))
		}

		// We check that the result contains all expected elements
		if !containsAll(result, test.expected) {
			t.Errorf("DFSIterative(%d) = %v; expected %v", test.start, result, test.expected)
		}
	}
}
