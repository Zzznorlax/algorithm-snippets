package sorts_test

import (
	"algorithms/sorts"
	"math/rand"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{0, 0, 0, 0, 0}

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	sorts.InsertionSort(arr)

	t.Logf("Sorted array: %v", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("arr[%v] = %v > arr[%v] = %v, arr=%v", i, arr[i], i+1, arr[i+1], arr)
		}
	}

}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{0, 2, 1, 3, 9, 6}
		sorts.InsertionSort(arr)
	}
}
