package main

import "fmt"

func main() {
	ch1 := incrementor()
	chanSum := puller(ch1)
	for n := range chanSum {
		fmt.Println(n)
	}
}

func incrementor() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func puller(ch chan int) chan int {
	chanOut := make(chan int)
	var sum int
	go func() {
		for n := range ch {
			sum += n
		}
		chanOut <- sum
		close(chanOut)
	}()
	return chanOut
}
