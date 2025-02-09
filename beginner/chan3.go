package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go producer(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}

func producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Println("sending the number: ", i)
		time.Sleep(3 * time.Millisecond)
	}
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		num := <-ch
		fmt.Println("receiving the number: ", num)
		fmt.Println("the square of the number is ", num*num)
	}
}
