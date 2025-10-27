package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Start: %v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	indexToRemove := 3
	result := removeElement(slice, indexToRemove)

	fmt.Printf("Result: %v, len=%d, cap=%d\n",
		result, len(result), cap(result))
}

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}
