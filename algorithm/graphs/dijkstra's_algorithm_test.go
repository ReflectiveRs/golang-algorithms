package graphs

import (
	"errors"
	"math"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		graph         *GraphDijkstra
		start         int
		expected      []int
		expectedError error
	}{
		{
			graph: &GraphDijkstra{
				Vertices: 5,
				Edges: map[int][]EdgeDijkstra{
					0: {{To: 1, Weight: 10}, {To: 2, Weight: 3}},
					1: {{To: 2, Weight: 1}, {To: 3, Weight: 2}},
					2: {{To: 1, Weight: 4}, {To: 3, Weight: 8}, {To: 4, Weight: 2}},
					3: {{To: 4, Weight: 7}},
					4: {},
				},
			},
			start:    0,
			expected: []int{0, 7, 3, 9, 5}, // Expected distances from the top 0
		},
		{
			graph: &GraphDijkstra{
				Vertices: 3,
				Edges: map[int][]EdgeDijkstra{
					0: {{To: 1, Weight: 1}, {To: 2, Weight: 4}},
					1: {{To: 2, Weight: 2}},
					2: {},
				},
			},
			start:    0,
			expected: []int{0, 1, 3}, // Expected distances from the top 0
		},
		{
			graph: &GraphDijkstra{
				Vertices: 4,
				Edges: map[int][]EdgeDijkstra{
					0: {{To: 1, Weight: 2}},
					1: {{To: 2, Weight: 3}},
					2: {{To: 3, Weight: 1}},
					3: {},
				},
			},
			start:    0,
			expected: []int{0, 2, 5, 6}, // Expected distances from the top 0
		},
		{
			graph: &GraphDijkstra{
				Vertices: 3,
				Edges: map[int][]EdgeDijkstra{
					0: {{To: 1, Weight: 1}},
					1: {{To: 0, Weight: 1}},
					2: {},
				},
			},
			start:    2,
			expected: []int{math.MaxInt32, math.MaxInt32, 0}, // Vertex 2 is not available
		},
		{
			graph: &GraphDijkstra{
				Vertices: 3,
				Edges: map[int][]EdgeDijkstra{
					0: {{To: 1, Weight: 4}},
					1: {{To: 2, Weight: -5}}, // Negative edge
					2: {},
				},
			},
			start:         0,
			expected:      nil,
			expectedError: ErrWeightIsNegative,
		},
	}

	for _, test := range tests {
		t.Run("Test Dijkstra", func(t *testing.T) {
			result, err := test.graph.Dijkstra(test.start)
			if err != nil {
				if !errors.Is(err, test.expectedError) {
					t.Errorf("Error expected %s, got %s", test.expectedError, err)
				}
			}
			for i, distance := range result {
				if distance != test.expected[i] {
					t.Errorf("Expected distance to vertex %d to be %d, got %d", i, test.expected[i], distance)
				}
			}
		})
	}
}
