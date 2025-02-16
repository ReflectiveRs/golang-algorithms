package sorting

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// We assume that the current element is minimal.
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j // Find the index of the minimum element
			}
		}
		// We swap the current element with the found minimum
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}
