package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Search sends query to Google search and returns results.
func Search(ctx context.Context, query string) (string, error) {
	// Simulate an external API call
	time.Sleep(2 * time.Second)
	return "some results", nil
}

// handler handles /search?q=foo requests
func handler(w http.ResponseWriter, req *http.Request) {
	// Create a new context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()

	// Get query parameter from URL
	query := req.FormValue("q")

	// Call Search function with context and query
	results, err := Search(ctx, query)

	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write results to response
	w.Write([]byte(results))
}

func main() {
	http.HandleFunc("/search", handler)
	http.ListenAndServe(":8080", nil)
}
