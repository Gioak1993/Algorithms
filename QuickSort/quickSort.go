package main

import (
	"fmt"
)

func main () {

	x := []int{1,2,5,3,5,7,4,7,4,34,234,6,56,4,3,23,2,5,7,6,7,78,7,5,545,43,453,3,52,35,23,5}
	fmt.Print(QuickSort(x))
	fmt.Print(QuickSortDescending(x))

}

func QuickSort(m []int) []int {

	if len(m) < 2 {
		return m
	}else {
		pivot := m[0]
		lesser := []int{}
		greater := []int{}

		for _,x := range m[1:]{
			if x <= pivot{
				lesser = append(lesser, x)
			}else{
				greater = append(greater, x)
			}
		}
		return append(append(QuickSort(lesser), pivot), QuickSort(greater)...)
	}
}

func QuickSortDescending(m []int) []int {

	if len(m) < 2 {
		return m
	}else {
		pivot := m[0]
		lesser := []int{}
		greater := []int{}

		for _,x := range m[1:]{
			if x <= pivot{
				lesser = append(lesser, x)
			}else{
				greater = append(greater, x)
			}
		}
		return append(append(QuickSortDescending(greater), pivot), QuickSortDescending(lesser)...)
	}
}