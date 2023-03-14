package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var someStrings = make(map[string]string)

func main() {
	i := 0
	var tmp string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		str := strconv.Itoa(i)
		tmp = input.Text()
		if tmp == "quit" {
			break
		}
		someStrings[str] = tmp
		i++
	}
	ch := make(chan string)
	//fmt.Println("len", len(someStrings))
	//fmt.Println(len(ch))
	for i := 0; i < len(someStrings); i++ {
		go capitalize(someStrings[strconv.Itoa(i)], ch)
		go lower(someStrings[strconv.Itoa(i)], ch)
		go appender(someStrings[strconv.Itoa(i)], ch)
	}
	for i := 0; i < len(someStrings)*3; i++ { // show buffered results
		fmt.Print(<-ch, "\n")
	}
}
func capitalize(s string, ch chan string) {
	s = strings.ToUpper(s)
	ch <- s
}
func lower(s string, ch chan string) {
	s = strings.ToLower(s)
	ch <- s
}
func appender(s string, ch chan string) {
	s = s + " This is test "
	ch <- s
}
