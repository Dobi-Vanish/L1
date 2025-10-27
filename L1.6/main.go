package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("\n1. Natural completion:")
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Natural: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Natural: completed")
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("\n2. Done channel:")
	done := make(chan struct{})
	go func() {
		i := 0
		for {
			i++
			select {
			case <-done:
				fmt.Println("Channel: stopped")
				return
			default:
				fmt.Printf("Channel: working %d...\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(500 * time.Millisecond)
	close(done)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n3. Context:")
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		i := 0
		for {
			i++
			select {
			case <-ctx.Done():
				fmt.Println("Context: stopped")
				return
			default:
				fmt.Printf("Context: working %d...\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(500 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n4. Timeout in select:")
	go func() {
		i := 0
		timeout := time.After(800 * time.Millisecond)
		for {
			i++
			select {
			case <-timeout:
				fmt.Println("Timeout: time's up!")
				return
			default:
				fmt.Printf("Timeout: processing %d...\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1 * time.Second)

	fmt.Println("\n5. Atomic flag:")
	var stopFlag int32
	go func() {
		i := 0
		for atomic.LoadInt32(&stopFlag) == 0 {
			i++
			fmt.Printf("Atomic: running %d...\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Atomic: flag set - stopping")
	}()
	time.Sleep(500 * time.Millisecond)
	atomic.StoreInt32(&stopFlag, 1)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n6. WaitGroup:")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Printf("WaitGroup: %d\n", i)
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("WaitGroup: completed")
	}()
	wg.Wait()

	fmt.Println("\nAll demonstrations completed!")
}
