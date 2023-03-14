package main

import (
	"fmt"
	"time"
)

func plus(a, b int, c chan int) {
	resultplus := a + b
	c <- resultplus
}

func minus(a, b int, c chan int) {
	resultminus := a - b
	c <- resultminus
}

func multi(a, b int, c chan int) {
	resultmulti := a * b
	c <- resultmulti
}

func devi(a, b int, c chan int) {
	resultdevi := a / b
	c <- resultdevi
}
func exp(a, b int, c chan int) {
	resultexp := a ^ b
	c <- resultexp

}

func main() {
	start := time.Now()
	var a, b int
	fmt.Println("enter two numbers to add")
	fmt.Scan(&a, &b)

	numberOfFunc := 4                 // number of items of channel
	c := make(chan int, numberOfFunc) // create buffered channel
	func(c chan int) {                // lambda
		go plus(a, b, c)
		go minus(a, b, c)
		go multi(a, b, c)
		go devi(a, b, c)
		go exp(a, b, c)
	}(c)

	time.Sleep(200 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Println("time spent:", elapsed)

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
