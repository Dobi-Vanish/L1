package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Worker number can't be less than 0")
		os.Exit(1)
	}

	fmt.Printf("Launching %d workers\n", numWorkers)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	dataChannel := make(chan string, numWorkers*2)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg, ctx)
	}

	wg.Add(1)
	go dataProducer(dataChannel, &wg, ctx)

	go func() {
		sig := <-signalChan
		fmt.Printf("\nSignal has been received: %v.\n", sig)
		cancel()
	}()

	wg.Wait()
	close(dataChannel)
}

func dataProducer(ch chan<- string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	defer fmt.Println("dataProducer has been stopped")

	counter := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			message := fmt.Sprintf("Mock data number #%d, time: %s", counter,
				time.Now().Format("15:04:05"))
			select {
			case ch <- message:
				counter++
			case <-ctx.Done():
				return
			case <-time.After(100 * time.Millisecond):
				continue
			}
			select {
			case <-time.After(500 * time.Millisecond):
			case <-ctx.Done():
				return
			}
		}
	}
}

func worker(id int, ch <-chan string, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	defer fmt.Printf("Worker %d has been stoopped\n", id)

	fmt.Printf("Worker %d has been started\n", id)

	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-ch:
			if !ok {
				return
			}

			processingTime := time.Duration(100+(id*50)) * time.Millisecond

			select {
			case <-time.After(processingTime):
				fmt.Printf("Worker %d: %s\n", id, data)
			case <-ctx.Done():
				return
			}
		}
	}
}
