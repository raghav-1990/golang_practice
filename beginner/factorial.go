// calculating factorial
package main

import (
	"fmt"
)

// without channels and goroutines
// func main() {
// 	n := 5
// 	var prod = 1
// 	fmt.Println("prod is ", prod)
// 	for i := n; i > 0; i-- {
// 		fmt.Println("i is now ", i)
// 		prod *= i
// 		fmt.Println("prod is  now", prod)
// 	}
// 	fmt.Println(prod)
// }

// with channels and goroutines
func main() {
	fact := factorial(5)
	for n := range fact {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	ch := make(chan int)
	go func() {
		product := 1
		for i := n; i > 0; i-- {
			product *= i
		}
		ch <- product
		close(ch)
	}()
	return ch
}
