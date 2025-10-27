package main

import (
	"fmt"
)

func main() {
	testValues := []interface{}{
		42,
		"hello",
		true,
		make(chan int),
		make(chan string),
		3.14,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
	}

	for _, value := range testValues {
		detectType(value)
	}

}

func detectType(v interface{}) {
	switch v := v.(type) {
	case int:
		fmt.Printf("Value: %v -> Type: int\n", v)
	case string:
		fmt.Printf("Value: %v -> Type: string\n", v)
	case bool:
		fmt.Printf("Value: %v -> Type: bool\n", v)
	case chan int:
		fmt.Printf("Value: %v -> Type: chan int\n", v)
	case chan string:
		fmt.Printf("Value: %v -> Type: chan string\n", v)
	case chan bool:
		fmt.Printf("Value: %v -> Type: chan bool\n", v)
	default:
		fmt.Printf("Value: %v -> Type: %T\n", v, v)
	}
}
