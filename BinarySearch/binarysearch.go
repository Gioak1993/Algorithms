package main

import (
	"fmt"
)

func main () {

	B := []int{11, 22, 34, 45, 54, 66, 77, 88, 99, 101, 122, 134, 145, 156, 163, 174, 186, 198, 205}
	fmt.Println(BinarySearch(B, 186))

}

func BinarySearchHelper (A []int, target int, left int, right int) int {

	mid := (right+left)/2

	if target == A[mid] {
		return mid 
	}else if target < A[mid]{
		return BinarySearchHelper(A, target, left, mid-1)
	}else{
		return BinarySearchHelper(A, target, mid+1, right)
	}

}


func BinarySearch (A []int, target int) int  {

	if A[0] > target || A[len(A)-1] < target {
		return -1
	} else {
		return BinarySearchHelper(A, target, 0, len(A)-1)
	}
}