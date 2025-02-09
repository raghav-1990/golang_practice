// incrementor with channels

package main

import "fmt"

func main() {
	c1 := incrementor("foo: ")
	c2 := incrementor("bar: ")
	c3 := puller(c1)
	c4 := puller(c2)
	// q := make(chan int)
	// var w chan int
	// fmt.Println(q, w)
	fmt.Println("the grand total is: ", <-c3+<-c4)
}

func incrementor(s string) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println(s, i)
		}
		close(ch)
	}()
	return ch
}

func puller(ch chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range ch {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
