
/*
Write a program which can compute the factorial of a given numbers. The results should be printed in a comma-separated sequence on a single line.

Suppose the following input is supplied to the program: 8
Then, the output should be: 40320
*/


package main
import "fmt"

func main(){
    var input int
    _, err:= fmt.Scanln(&input)
    if err != nil {
        fmt.Printf("error scanning the input number: %v", err)
    }
    result, factorialError := factorial(input) 
    if factorialError != nil {
        fmt.Printf("error calculating the factorial: %v", factorialError)
    }
    fmt.Printf("the factorial of %v is %v", input, result)
}

func factorial(input int) (uint64, error) {
    var factorial uint64 = 1
if input < 0 {
    return 0, fmt.Errorf("Provided incorrect number, less than 0")
}
if input == 0 {
    return 1, nil
}

for i := 1; i <= input; i++ {
    factorial *= uint64(i)
}
return factorial, nil
}
