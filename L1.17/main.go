package main

import "fmt"

func main() {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	targets := []int{7, 1, 19, 8, 0, 20}

	fmt.Printf("Provided: %v\n\n", sortedArray)

	for _, target := range targets {
		index := binarySearch(sortedArray, target)
		if index != -1 {
			fmt.Printf("Number %d at position %d\n", target, index)
		} else {
			fmt.Printf("Number %d has not been founded\n", target)
		}
	}
}

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
