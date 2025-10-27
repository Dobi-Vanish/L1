package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Worker number can't be less than 0")
		return
	}

	fmt.Printf("Launching %d workers\n", numWorkers)

	dataChannel := make(chan string)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	go dataProducer(dataChannel)

	wg.Wait()
}

func dataProducer(ch chan<- string) {
	counter := 1
	for {
		message := fmt.Sprintf("Mock data number #%d, time: %s", counter, time.Now().Format("15:04:05"))

		ch <- message

		counter++

		time.Sleep(500 * time.Millisecond)
	}
}

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d has been started\n", id)

	for data := range ch {
		fmt.Printf("Worker %d: %s\n", id, data)
	}

	fmt.Printf("Worker %d finished\n", id)
}
