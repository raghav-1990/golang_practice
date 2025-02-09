package main

import (
	"fmt"
)

func main() {
	n := 3
	ch := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				ch <- j
			}
			done <- true
			// time.Sleep(30 * time.Millisecond)
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
