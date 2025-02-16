package search

func LinearSearch(arr []int, target int) int {
	for index, value := range arr {
		if value == target {
			return index // Return the index of the found element
		}
	}
	return -1 // Element not found
}
