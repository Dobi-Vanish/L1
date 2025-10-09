package main

import "fmt"

func main() {
	a, b := 4, 2
	fmt.Printf("a = %d, b = %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a = %d, b = %d\n", a, b)
}
