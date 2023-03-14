package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Search sends a query to Google search and returns results.
func Search(ctx context.Context, query string) ([]string, error) {
	// Create a new HTTP client with timeout
	client := &http.Client{Timeout: 5 * time.Second}

	// Prepare the Google search URL with query
	searchURL := fmt.Sprintf("https://www.google.com/search?q=%s", query)

	// Create a new HTTP request with the prepared URL
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}

	// Set user agent header to mimic a browser
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	// Send the request and parse the response
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Use goquery to parse the HTML response
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Extract the search result titles and links from the HTML response
	var results []string
	doc.Find(".rc .r a").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		link, _ := s.Attr("href")
		if link != "#" {
			results = append(results, fmt.Sprintf("%s (%s)", title, link))
		}
	})

	return results, nil
}

// handler handles /search?q=foo requests
func handler(w http.ResponseWriter, req *http.Request) {
	// Create a new context with a timeout of 5 seconds
	ctx, cancel := context.WithTimeout(req.Context(), 5*time.Second)
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
	for _, result := range results {
		fmt.Fprintf(w, "%s\n", result)
	}
}

func main() {
	http.HandleFunc("/search", handler)
	http.ListenAndServe(":8080", nil)
}
