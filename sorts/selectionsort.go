package sorts

// SelectionSort implements selection sort
func SelectionSort(arr []int, start int, end int) []int {

	for i := start; i < end; i++ {
		currentElement := arr[i]

		minElementIndex := i
		for j := i + 1; j <= end; j++ {
			if arr[j] < arr[minElementIndex] {
				minElementIndex = j
			}
		}

		arr[i] = arr[minElementIndex]
		arr[minElementIndex] = currentElement
	}

	return arr
}
