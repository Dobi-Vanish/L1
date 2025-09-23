package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}

	calculateSquaresWithWaitGroup(numbers)
}

func calculateSquaresWithWaitGroup(numbers [5]int) {
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)

		go func(x int) {
			defer wg.Done()
			square := x * x
			fmt.Printf("pow(%d) = %d\n", x, square)
		}(num)
	}

	wg.Wait()
}
