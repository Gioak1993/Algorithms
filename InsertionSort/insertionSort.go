package main

import (
	"fmt"
)

func main() {

	array := []int{3, 7, 1, 0, 5, 6, 45, 78, 20, 45, 6, 7, 3333, 223, 5, 7, 9, 0, 333, 2, 5, 7, 44, 75433, 677, 13}
	fmt.Print(insertionSortIncresing(array))
	fmt.Print(insertionSortDecreasing(array))

}

func insertionSortIncresing(A []int) []int {

	for j := 1; j < len(A); j++ {
		key := A[j]
		// insert A[j] into the sorted sequence
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	return A
}

func insertionSortDecreasing(A []int) []int {

	for j := 1; j < len(A); j++ {
		key := A[j]
		i := j - 1
		for i >= 0 && A[i] < key {
			A[i+1] = A[i]
			i = i - 1
		}
		A[i+1] = key
	}
	return A
}
