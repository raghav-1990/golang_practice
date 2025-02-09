package main

import (
	"fmt"
	"strconv"
	"strings"
)

var bigString interface {
	ReadString(string)
	PrintString(string)
}

func ReadString(input string) {
	
}
func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error occurred while reading input: %v\n", err)
	}
	numbers := strings.Split(input, ",")
	var num = make([]int, len(numbers))
	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	fmt.Println(num)
}
