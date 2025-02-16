package sorting

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // Flag to track if exchanges have been made
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Let's swap places arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If there were no exchanges during the pass, then the array is sorted
		if !swapped {
			break
		}
	}
}
