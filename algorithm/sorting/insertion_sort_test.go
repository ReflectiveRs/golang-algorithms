package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "check with typical array",
			input:    []int{64, 34, 25, 12, 22, 11, 90},
			expected: []int{11, 12, 22, 25, 34, 64, 90},
		},
		{
			name:     "check, array is already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "check, array is already sorted",
			input:    []int{-1, -2, -3, -4, -5},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "check, array sorted in reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "check, empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "check, array with one element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, test := range tests {
		// Copy the input array so as not to change it
		input := make([]int, len(test.input))
		copy(input, test.input)

		InsertionSort(input)

		if !reflect.DeepEqual(input, test.expected) {
			t.Errorf("%s, InsertionSort(%v) = %v; want %v", test.name, test.input, input, test.expected)
		}
	}
}
