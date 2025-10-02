package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("\nSafeMap with Mutex:")
	testSafeMap()

	fmt.Println("\nRace condition:")
	raceTest()
}

func testSafeMap() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("worker%d_key%d", id, j)
				safeMap.Set(key, id*100+j)
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("Total entries: %d\n", safeMap.Len())

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("worker%d_key%d", id, j)
				if value, exists := safeMap.Get(key); exists {
					fmt.Printf("Read: %s = %d\n", key, value)
				}
			}
		}(i)
	}

	wg.Wait()
}

func raceTest() {
	unsafeMap := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)

				mu.Lock()
				unsafeMap[key] = id + j
				mu.Unlock()
			}
		}(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("key_%d_%d", id, j)

				mu.Lock()
				_ = unsafeMap[key]
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Race test completed successfully")
}

type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	value, exists := sm.data[key]
	return value, exists
}

func (sm *SafeMap) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.data)
}
