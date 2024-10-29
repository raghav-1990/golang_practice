/* Write a program which accepts a sequence of comma-separated numbers from console and generate an slice out of them. Return the slice.

Suppose the following input is supplied to the program: 34, 67, 55, 33, 12, 98.

Then, the output should be: [34 67 55 33 12 98]
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
