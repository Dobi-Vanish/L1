package main

import (
    "fmt"
    "time"
)

func Sleep(duration time.Duration) {
    <-time.After(duration)
}

func main() {
    fmt.Println("Start:", time.Now())
    Sleep(2 * time.Second)
    fmt.Println("2 seconds have passed:", time.Now())
}