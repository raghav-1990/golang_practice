package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(&wg)
	go bar(&wg)
	wg.Wait()
}

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("This is foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
}

func bar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("This is bar: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
}
