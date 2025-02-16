package graphs

import (
	"math"
	"testing"
)

func TestFloydWarshall(t *testing.T) {
	tests := []struct {
		graph    GraphFloydWarshall
		expected [][]int
	}{
		{
			graph: GraphFloydWarshall{
				Graph: [][]int{
					{0, 3, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 7, 0},
					{0, 0, 0, 0, 2},
					{6, 0, 0, 0, 0},
				},
			},
			expected: [][]int{
				{0, 3, 4, 11, 13},
				{16, 0, 1, 8, 10},
				{15, 18, 0, 7, 9},
				{8, 11, 12, 0, 2},
				{6, 9, 10, 17, 0},
			},
		},
		{
			graph: GraphFloydWarshall{
				Graph: [][]int{
					{0, 1, 4, 0},
					{0, 0, 2, 0},
					{0, 0, 0, 1},
					{0, 0, 0, 0},
				},
			},
			expected: [][]int{
				{0, 1, 3, 4},
				{math.MaxInt32, 0, 2, 3},
				{math.MaxInt32, math.MaxInt32, 0, 1},
				{math.MaxInt32, math.MaxInt32, math.MaxInt32, 0},
			},
		},
		{
			graph: GraphFloydWarshall{
				Graph: [][]int{
					{0, -1, 4, 0},
					{0, 0, 3, 2},
					{0, 0, 0, -1},
					{0, 0, 0, 0},
				},
			},
			expected: [][]int{
				{0, -1, 2, 1},
				{math.MaxInt32, 0, 3, 2},
				{math.MaxInt32, math.MaxInt32, 0, -1},
				{math.MaxInt32, math.MaxInt32, math.MaxInt32, 0},
			},
		},
	}

	for _, test := range tests {
		result := FloydWarshall(test.graph)
		for i := range test.expected {
			for j := range test.expected[i] {
				if result[i][j] != test.expected[i][j] {
					t.Errorf("Test failed for input %v. Expected %v but got %v", test.graph, test.expected, result)
				}
			}
		}
	}
}
