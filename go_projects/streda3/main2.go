



/*
package main

import (
	"fmt"
	"strconv"
	"strings"
)

var someStrings = make(map[string]string)

func main() {
	i := 0
	var tmp string
	for {
		str := strconv.Itoa(i)
		fmt.Scanf("%v", &tmp)
		if tmp == "quit" {
			break
		}
		fmt.Println("tmp is", tmp)
		someStrings[str] = tmp //"asdfg"
		i++
	}
	ch := make(chan string, len(someStrings))
	go capitalize(someStrings["1"], ch)
	go lower(someStrings["1"], ch)
	go appender(someStrings["1"], ch)
	//fmt.Println(len(ch))
	for i := 0; i < len(someStrings); i++ {
		fmt.Println(<-ch)
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
	s = s + " This is test"
	ch <- s
}
*/
