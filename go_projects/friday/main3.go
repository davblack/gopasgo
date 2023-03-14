package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SearchResult represents a single search result
type SearchResult struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// SearchService is a local search service that returns search results
type SearchService struct{}

// Search sends query to a local search service and returns results.
func (s *SearchService) Search(ctx context.Context, query string) ([]SearchResult, error) {
	// Simulate a search service that returns fake search results
	// In a real implementation, this function would call an actual search service
	// and return the search results returned by the service.
	results := []SearchResult{
		{Title: "Search Result 1", Link: "http://example.com/result1"},
		{Title: "Search Result 2", Link: "http://example.com/result2"},
		{Title: "Search Result 3", Link: "http://example.com/result3"},
	}

	return results, nil
}

// handler handles /search?q=foo requests
func handler(w http.ResponseWriter, req *http.Request) {
	// Create a new context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
	defer cancel()

	// Get query parameter from URL
	query := req.FormValue("q")

	// Create a new search service and call the Search method with context and query
	service := &SearchService{}
	results, err := service.Search(ctx, query)

	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write results to response
	jsonResults, err := json.Marshal(results)
	if err != nil {
		fmt.Println("error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResults)
}

func main() {
	http.HandleFunc("/search", handler)
	http.ListenAndServe(":8080", nil)
}
