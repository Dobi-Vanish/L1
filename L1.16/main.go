package main

import "fmt"

func main() {
	arr := []int{9, -3, 5, 2, 6, 8, -6, 1, 3}
	fmt.Printf("Provided: %v\n", arr)

	sorted := quickSort(arr)
	fmt.Printf("Result: %v\n", sorted)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[len(arr)/2]

	var less, equal, greater []int

	for _, num := range arr {
		switch {
		case num < pivot:
			less = append(less, num)
		case num == pivot:
			equal = append(equal, num)
		case num > pivot:
			greater = append(greater, num)
		}
	}

	return append(append(quickSort(less), equal...), quickSort(greater)...)
}
