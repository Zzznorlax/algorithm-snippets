package sorts

// MergeSort implements merge sort
func MergeSort(arr []int, start int, end int) []int {

	if end-start < 1 {
		return arr
	}

	mid := (start + end) / 2

	MergeSort(arr, start, mid)
	MergeSort(arr, mid+1, end)

	return merge(arr, start, mid, end)
}

func merge(arr []int, left int, mid int, right int) []int {

	leftIndex := left
	rightIndex := mid + 1

	combined := []int{}

	for leftIndex <= mid && rightIndex <= right {
		if arr[leftIndex] < arr[rightIndex] {
			combined = append(combined, arr[leftIndex])
			leftIndex++
		} else {
			combined = append(combined, arr[rightIndex])
			rightIndex++
		}
	}

	start := mid + 1
	end := right
	if leftIndex != mid {
		start = leftIndex
		end = mid
	}

	for i := start; i <= end; i++ {
		combined = append(combined, arr[i])
	}

	for i, j := left, 0; j < len(combined); i, j = i+1, j+1 {
		arr[i] = combined[j]
	}

	return arr
}
