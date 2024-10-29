/*
With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value
*/

package main

import "fmt"

func main() {
	var num int

	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Printf("Error occurred while reading input: %v\n", err)
	}
	squareMap := make(map[int]int)

	for i := 1; i <= num; i++ {
		squareMap[i] = i * i
	}
	fmt.Println(" the map is: ", squareMap)
}
