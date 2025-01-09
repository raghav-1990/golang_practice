// Implement a producer-consumer problem in Go using buffered channels, where one Goroutine produces numbers, and another consumes them to calculate their squares.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go Producer(ch, &wg)
	go Consumer(ch, &wg)
	wg.Wait()
}

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:= 0; i < 10; i++ {
		fmt.Printf("Number produced is : %v \n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:= 0; i < 10; i++ {
		value := <- ch 
		fmt.Printf("Number comsumed is %v and square of the number is %v \n", value, value*value)
	}
}
