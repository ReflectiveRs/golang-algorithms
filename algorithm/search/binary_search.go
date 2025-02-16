package search

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Return the index of the found element
		} else if arr[mid] < target {
			left = mid + 1 // We are looking in the right half
		} else {
			right = mid - 1 // We are looking in the left half
		}
	}

	return -1 // Element not found
}
