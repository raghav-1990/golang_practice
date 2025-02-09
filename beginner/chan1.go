package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go printOdd(ch, &wg)
	go printEven(ch, &wg)
	wg.Wait()
}

func printOdd(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- true
		fmt.Println("This is odd : ", i)
		<-ch
	}
}

func printEven(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		<-ch
		fmt.Println("This is even : ", i)
		ch <- true
	}
}
