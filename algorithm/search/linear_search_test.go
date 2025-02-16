package search

import "testing"

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{arr: []int{5, 3, 8, 4, 2}, target: 4, expected: 3},
		{arr: []int{1, 2, 3, 4, 5}, target: 1, expected: 0},
		{arr: []int{1, 2, 3, 4, 5}, target: 5, expected: 4},
		{arr: []int{10, 20, 30, 40, 50}, target: 25, expected: -1},
		{arr: []int{-1, -3, -4, -5}, target: -4, expected: 2},
		{arr: []int{}, target: 1, expected: -1},
		{arr: []int{10}, target: 10, expected: 0},
	}

	for _, test := range tests {
		result := LinearSearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("LinearSearch(%v, %d) = %d; expected %d", test.arr, test.target, result, test.expected)
		}
	}
}
