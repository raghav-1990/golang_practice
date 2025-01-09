//Create a program that launches two goroutines: one to print even numbers and another to print odd numbers, up to 10. Use channels to coordinate between the two goroutines.

package main

import (
	"fmt"
	"sync"
)

func main(){
	ch := make(chan bool)
	var wg sync.WaitGroup

	wg.Add(2)
	go printEven(ch, &wg)
	go printOdd(ch, &wg)

	wg.Wait()
}

func printEven(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:= 0; i < 10; i++{
		if i%2 == 0 {
			ch <- true
			fmt.Println("this is even:- ", i)
			<- ch
		}
	}
}

func printOdd(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i:= 1; i < 10; i++{
		if i%2 != 0 {
			<- ch
			fmt.Println("this is odd:- ", i)
			ch <- true
		}
	}
}
