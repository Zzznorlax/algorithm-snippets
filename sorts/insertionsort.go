package sorts

// InsertionSort implements insertion sort
func InsertionSort(array []int) []int {

	for i := 1; i < len(array); i++ {

		element := array[i]

		j := i - 1
		for ; j >= 0 && array[j] > element; j-- {
			array[j+1] = array[j]
		}

		array[j+1] = element
	}

	return array
}
