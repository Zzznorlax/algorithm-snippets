package sorts_test

import (
	"algorithms/sorts"
	"math/rand"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	arrLength := 10
	arr := []int{}

	for i := 0; i < arrLength; i++ {
		arr = append(arr, rand.Intn(100))
	}

	start := 0
	end := len(arr) - 1

	sorts.SelectionSort(arr, start, end)

	t.Logf("Sorted array: %v", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("arr[%v] = %v > arr[%v] = %v, arr=%v", i-1, arr[i-1], i, arr[i], arr)
		}
	}

}

func BenchmarkSelectionSort(b *testing.B) {

	arrLength := 10

	arr := []int{}

	for i := 0; i < arrLength; i++ {
		arr = append(arr, rand.Intn(100))
	}

	arrCopy := []int{}

	start := 0
	end := len(arrCopy) - 1

	for i := 0; i < b.N; i++ {

		copy(arrCopy, arr)

		sorts.SelectionSort(arrCopy, start, end)
	}
}
