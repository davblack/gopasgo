package main

import "fmt"

//In this example, add takes two parameters of type int.
func add(x int, y int) int {
	return x + y
}

//When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func adds(x, y int) int {
	return x + y
}

//A function can return any number of results. The swap function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

//Go's return values may be named. If so, they are treated as variables defined at the top of the function.
//These names should be used to document the meaning of the return values.
//A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(adds(99, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
