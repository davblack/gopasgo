package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// Create an unbuffered channel variable
	ch := make(chan int, 3)
	// Create a mutex variable
	var mu sync.Mutex
	// Launch 3 goroutines
	wg.Add(1)
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 1")
		mu.Unlock() // Unlock the mutex
		ch <- 1     // Send 1 to the channel
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 2")
		mu.Unlock() // Unlock the mutex
		ch <- 2     // Send 2 to the channel
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 3")
		mu.Unlock() // Unlock the mutex
		ch <- 3     // Send 3 to the channel
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Main function")
	// Receive values from the channel in order using a for loop with counter
	for i := 0; i < 3; i++ {
		j := <-ch            // Receive j from chanel
		fmt.Printf("%d ", j) // Print j
	}
}

/*package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create an unbuffered channel variable
	ch := make(chan int)

	// Create a mutex variable
	var mu sync.Mutex

	// Launch 3 goroutines
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 1")
		mu.Unlock() // Unlock the mutex
		ch <- 1     // Send 1 to the channel
	}()
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 2")
		mu.Unlock() // Unlock the mutex
		ch <- 2     // Send 2 to the channel
	}()
	go func() {
		mu.Lock() // Lock the mutex
		fmt.Println("Goroutine 3")
		mu.Unlock() // Unlock the mutex
		ch <- 3     // Send 3 to the channel
	}()

	fmt.Println("Main function")

	// Receive values from the channel in order using a for loop with counter
	for i := 0; i < 3; i++ {
		j := <-ch            // Receive j from chanel
		fmt.Printf("%d ", j) // Print j
	}
}
*/
