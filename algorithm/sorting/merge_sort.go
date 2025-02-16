package sorting

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// Compare the elements from the left and right halves and add the smaller element to the result
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Add the remaining elements from the left half
	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	// Add the remaining elements from the right half
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
