package main

import (
	"fmt"
	"sync"
)

var count int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go counter(&wg)
	}
	wg.Wait()
	fmt.Println("the final counter value is ", count)
}

func counter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		mu.Lock()
		count += 1
		mu.Unlock()
	}
}
