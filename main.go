package main

import (
	"algorithms/sorts"
	"fmt"
)

func main() {

	arr := []int{0, 7, 6, 2, 5, 9}

	sorts.InsertionSort(arr)

	fmt.Println(arr)
}
