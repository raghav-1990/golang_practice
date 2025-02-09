package main

import "fmt"

func main() {
	ans := divide(10, 0)
	fmt.Println(ans)

	fmt.Println("Program continues after recovered panic")
}

func divide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered!")
		}
	}()

	if b == 0 {
		fmt.Println("reached here")
		panic("Divide by Zero operation!!")
	}
	fmt.Println("reached here too")
	return (a / b)
}
