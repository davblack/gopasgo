package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
func calc(a, b int) int {

	return a + b
}

func main() {
	var a, b int
	fmt.Println("enter two numbers to add")
	fmt.Scan(&a, &b)
	fmt.Println(calc(a, b))
}
*/

func main() {
	if len(os.Args) == 4 { // check how many parameters are input
		a, err := strconv.Atoi(os.Args[1]) //converts the input arg "a" to integer
		if err != nil {                    // if there is error, prints an error msg
			fmt.Println("Error:", err, ", please enter an integer instead")
			return // exits the program because of error
		}
		b, err := strconv.Atoi(os.Args[3]) //converts the input arg "b" to integer
		if err != nil {
			fmt.Println("Error:", err, ", please enter an integer instead")
			return
		}
		switch os.Args[2] {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "/":
			if b == 0 {
				fmt.Println("Error:", "you can't divide by 0")
				return
			}
			fmt.Println(a / b)
		case ".":
			fmt.Println(a * b)
		default:
			fmt.Println("Error:", "please use '+'  '-'  '.'  '/' as the last parameter")
		}
	} else {
		// if the previous "if" statement is not true prints an error
		fmt.Println("Error:", "Please give two integer and a basic math symbol as input, example: '2 + 7'")
	}
}
