package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

}

// the func below causes deadlock. because there is no go routine to pick up the channel value/
// the func above eases the deadloack because there are 2 separate goroutines
// func main() {
// 	ch := make(chan int)
// 	ch <- 1
// 	fmt.Println(<-ch)
// }
