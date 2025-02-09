package main

import "fmt"

func main() {
	mainSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	firstSlice := mainSlice[:len(mainSlice)/2]
	secondSlice := mainSlice[len(mainSlice)/2:]
	ch := make(chan int)
	go calc(ch, firstSlice)
	go calc(ch, secondSlice)

	fmt.Println("ch b4 sum1 is: ", ch)
	sum1 := <-ch
	fmt.Println("ch after sum1 is: ", ch)
	fmt.Println("ch b4 sum2 is: ", ch)
	sum2 := <-ch
	fmt.Println("ch after sum2 is: ", ch)
	fmt.Println("the sum is: ", sum1+sum2)
}

func calc(ch chan int, s []int) {
	var result int
	fmt.Println("ch b4 is: ", ch)
	for _, val := range s {
		result += val
	}
	ch <- result
	fmt.Println("ch after is: ", ch)

}
