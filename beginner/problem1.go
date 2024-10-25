/*
Exercise 001:

Write a program which will find all such numbers which are divisible by 7 but are not a multiple of 5, between 2000 and 3200 (both included).  The numbers obtained should be printed in a comma-separated sequence on a single line.
*/


package main
import "fmt"
import "strconv"
import "strings"

func main() {
  numbers := divisible(2000,3200)
  fmt.Println(numbers)
}

func divisible(startingNumber, endingNumber int) string {
    var nums []string
    for i:= startingNumber; i <= endingNumber; i++ {
        if i%7 == 0 && i%5 != 0 {
            nums = append(nums, strconv.Itoa(i))
        }
    }
    return strings.Join(nums, ", ")
}
