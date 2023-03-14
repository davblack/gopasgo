package main

import (
	"bufio"   // to scan/import my .txt file
	"fmt"     // i need to print out
	"log"     // to log
	"os"      // i need to open files
	"strings" // to work with text
)

func main() {
	//open a file
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//creating a scanner to read the file
	scanner := bufio.NewScanner(file)

	//input the world to scan for
	wordsToCount := []string{"her"}

	//making a map to store the counted words
	wordCounts := make(map[string]int)

	//scan all line in the txt file
	for scanner.Scan() {
		//splitting the lines to words
		words := strings.Fields(scanner.Text())
		//loop over each word
		for _, word := range words {
			//making lowercase every word in txt file
			word = strings.ToLower(word)
			//check if the word is in the list of words to count
			for _, countWord := range wordsToCount {
				if word == countWord {
					// Increment the count for this word
					wordCounts[countWord]++
				}
			}
		}
	}
	// Print out the word counts
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
