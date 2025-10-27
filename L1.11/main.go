package main

import "fmt"

func main() {
	A := []int{1, 2, 2, 3, 4, 4, 5}
	B := []int{2, 3, 3, 4, 4, 6, 7}

	intersection := IntersectUnique(A, B)
	fmt.Printf("A = %v\n", A)
	fmt.Printf("B = %v\n", B)
	fmt.Printf("Пересечение = %v\n", intersection)
}

func IntersectUnique(a, b []int) []int {
	setA := make(map[int]bool)
	for _, item := range a {
		setA[item] = true
	}

	setB := make(map[int]bool)
	for _, item := range b {
		setB[item] = true
	}

	var result []int
	for item := range setA {
		if setB[item] {
			result = append(result, item)
		}
	}

	return result
}
