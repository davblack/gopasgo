package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// makes a HTTP GET request to google.com
	resp, err := http.Get("http://google.com/")
	if err != nil { // if there is an error exits the program
		log.Fatalf("Encountered the following error", err)
	}
	// defers the closing of the response body until after the function had finished executing
	defer resp.Body.Close()

	// reads the response body into a byte slice called body using ioutil.ReadAll.
	body, err := ioutil.ReadAll(resp.Body)

	// prints the response body as a string to the console using fmt.Println.
	fmt.Println(string(body))
}
