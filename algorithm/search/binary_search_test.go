package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr      []int
		target   int
		expected int
	}{
		{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, target: 3, expected: 2},
		{arr: []int{10, 20, 30, 40, 50}, target: 10, expected: 0},
		{arr: []int{-1, -2, -3, -4, -5, -6, -7, -8}, target: -4, expected: 3},
		{arr: []int{1, 2, 3, 4, 5}, target: 7, expected: -1},
		{arr: []int{}, target: 1, expected: -1},
	}

	for _, test := range tests {
		result := BinarySearch(test.arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d; expected %d", test.arr, test.target, result, test.expected)
		}
	}
}
