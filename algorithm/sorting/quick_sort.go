package sorting

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2] // Selecting reference element
	var left []int
	var right []int
	var middle []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v) // Elements are smaller than the reference
		} else if v > pivot {
			right = append(right, v) // Elements are larger than the reference
		} else {
			middle = append(middle, v) // Elements equal to the reference
		}
	}

	// Recursively sort the subarrays and merge them
	return append(append(QuickSort(left), middle...), QuickSort(right)...)
}
