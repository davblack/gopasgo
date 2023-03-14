package main

import (
	"fmt"
	//"io/ioutil"
	//"log"
)

var user_map = map[string]string{
	"david":  "pass123",
	"radko":  "123pass",
	"martin": "password",
	"ivan":   "pass1234",
}

func main() {
	//var username, pass string
	var pass string
	var username string

	fmt.Println("please enter yo name fam:")
	fmt.Scanln(&username)

	fmt.Println("please enter yo pw fam:")
	fmt.Scanln(&pass)

	contain, err := user_map[username]
	if err != false && contain == pass {
		fmt.Println("you logged")
	} else {
		fmt.Println("wrong acc")
	}

}
