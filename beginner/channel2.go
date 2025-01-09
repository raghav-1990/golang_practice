// Write a Go program to compute the sum of numbers in a slice concurrently by dividing the slice into two parts and processing each part in a separate Goroutine. Combine the results using a channel.

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	m := len(nums) / 2
	firstHalf := nums[:m]
	secondHalf := nums[m:]

	go sum(firstHalf, ch)
	go sum(secondHalf, ch)

	firstSum := <-ch
	secondSum := <-ch

	fmt.Println("the total sum is : ", firstSum+secondSum)
}

func sum(nums []int, ch chan int) {
	var result int
	for _, value := range nums {
		result += value
	}
	ch <- result
}

