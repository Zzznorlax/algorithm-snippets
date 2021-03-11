package sorts

// QuickSort implements quick sort
func QuickSort(arr []int, start int, end int) []int {
	if start >= end {
		return arr
	}

	pivotIndex := partition(arr, start, end)

	QuickSort(arr, start, pivotIndex-1)
	QuickSort(arr, pivotIndex+1, end)

	return arr
}

func partition(arr []int, start int, end int) int {
	pivot := arr[end]

	i := start - 1
	for j := start; j <= end; j++ {
		element := arr[j]

		if arr[j] <= pivot {
			i++
			arr[j] = arr[i]
			arr[i] = element

		}
	}

	return i
}
