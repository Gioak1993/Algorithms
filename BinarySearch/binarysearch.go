package main

import (
	"fmt"
)

func main () {

	B := []int{11, 22, 34, 45, 54, 66, 77, 88, 99, 101, 122, 134, 145, 156, 163, 174, 186, 198, 205}
	fmt.Println(BinarySearch(B, 54))

}

func BinarySearchHelper (A []int, target int, index int) string {

	right := len(A)
	mid := right / 2
	fmt.Println(A)

	if A[0] > target || A[len(A)-1] < target {
		return fmt.Sprint("The target is not on the list")
	} else if A[mid] == target {
		index += mid
		return fmt.Sprintf("The target %d was found at index: %d", A[mid], index)
	} else if A[mid] > target {
		A = A[:mid]
		return BinarySearchHelper(A, target, index)
	} else if A[mid] < target{
		index += mid
		A = A[mid:]
		return BinarySearchHelper(A, target, index)
	} else {
		return fmt.Sprint("The target was found not found")
	}
}


func BinarySearch (A []int, target int) string {

	if A[0] > target || A[len(A)-1] < target {
		return fmt.Sprint("The target is not on the list")
	} else {
		return BinarySearchHelper(A, target, 0)
	}
}