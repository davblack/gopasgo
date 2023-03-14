/*package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	// Create a context.Background value as the root context
	ctx := context.Background()

	// Create a channel to synchronize the goroutines
	done := make(chan bool)

	// Define a string to be processed by the goroutines
	s := "hello world"

	// Launch the first goroutine with a context.WithValue
	go func(ctx context.Context) {
		fmt.Println("First goroutine:", strings.ToUpper(s))
		done <- true // Send a signal when done
	}(context.WithValue(ctx, "key", "value"))

	// Wait for the first goroutine to finish
	<-done

	// Launch the second goroutine with a context.WithCancel
	ctx2, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		fmt.Println("Second goroutine:", strings.ToLower(s))
		cancel()     // Cancel the context when done
		done <- true // Send a signal when done
	}(ctx2)

	// Wait for the second goroutine to finish
	<-done

	// Launch the third goroutine with a context.WithDeadline
	deadline := time.Now().Add(10 * time.Second) // Set a deadline 10 seconds from now
	ctx3, cancel := context.WithDeadline(ctx, deadline)
	go func(ctx context.Context) {
		fmt.Println("Third goroutine:", s+" this is context")
		cancel()     // Cancel the context when done
		done <- true // Send a signal when done
	}(ctx3)

	// Wait for all goroutines to finish
	<-done

}
*/