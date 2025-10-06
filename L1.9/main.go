package main

import (
	"fmt"
)

func main() {

	stage1 := make(chan int)
	stage2 := make(chan int)

	data := []int{2, 4, 6, 8, 10, 12, 14, 15, 16}

	go producer(data, stage1)
	go multiplier(stage1, stage2)
	consumer(stage2)

	fmt.Println("Конвейер завершен")
}

func producer(numbers []int, out chan<- int) {
	defer close(out)

	for _, num := range numbers {
		fmt.Printf("Producer: sent number %d\n", num)
		out <- num
	}
}

func multiplier(in <-chan int, out chan<- int) {
	defer close(out)

	for num := range in {
		result := num * 2
		fmt.Printf("Multyplier: numbers multyplied %d → %d\n", num, result)
		out <- result
	}
}

func consumer(in <-chan int) {
	for result := range in {
		fmt.Printf("Result: %d\n", result)
	}
}
