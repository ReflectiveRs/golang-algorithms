package graphs

import "testing"

func TestBFS(t *testing.T) {
	graph := NewGraphBFS()
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
		{0, []int{0, 1, 2, 3, 4, 5, 6}}, // Expected order
		{1, []int{1, 3, 4}},             // Expected order
		{2, []int{2, 5, 6}},             // Expected order
	}

	for _, test := range tests {
		result := graph.BFS(test.start)

		// We check that the length of the result matches the expected one
		if len(result) != len(test.expected) {
			t.Errorf("BFS(%d) = %v; expected length %d", test.start, result, len(test.expected))
		}

		// We check that the result contains all expected elements
		if !containsAll(result, test.expected) {
			t.Errorf("BFS(%d) = %v; expected %v", test.start, result, test.expected)
		}
	}
}
