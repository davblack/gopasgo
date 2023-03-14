package main

import "fmt"

//Go has only one looping construct, the for loop.

/*
The basic for loop has three components separated by semicolons:
	the init statement: executed before the first iteration
	the condition expression: evaluated before every iteration
	the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration,
and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false
*/
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

	//The init and post statements are optional.
	//At that point you can drop the semicolons: C's while is spelled for in Go.
	summ := 1
	for summ < 1000 {
		summ += summ
	}

	fmt.Println(summ)
}
