package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	workTime := 5

	if len(os.Args) > 1 {
		if sec, err := strconv.Atoi(os.Args[1]); err == nil && sec > 0 {
			workTime = sec
		} else {
			fmt.Println("Error: incorrect type of launching number")
			os.Exit(1)
		}
	}

	done := make(chan struct{})
	dataChannel := make(chan int)

	go sender(dataChannel, done)
	go receiver(dataChannel, done)

	timer := time.NewTimer(time.Duration(workTime) * time.Second)

	<-timer.C
	fmt.Printf("\nTime is up...\n")

	close(done)

	time.Sleep(100 * time.Millisecond)
}

func sender(ch chan<- int, done <-chan struct{}) {
	counter := 1
	for {
		select {
		case <-done:
			fmt.Println("Sender done")
			return
		case <-time.After(300 * time.Millisecond):
			select {
			case ch <- counter:
				fmt.Printf("Sent: %d\n", counter)
				counter++
			case <-done:
				return
			}
		}
	}
}

func receiver(ch <-chan int, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Receiver done")
			return
		case value := <-ch:
			fmt.Printf("Received: %d\n", value)
		}
	}
}
