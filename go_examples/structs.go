package main

import "fmt"

//A struct is a collection of fields.

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	p := &v //Struct fields can be accessed through a struct pointer.
	v.X = 4 //Struct fields are accessed using a dot.
	p.X = 1e9
	fmt.Println(v.X)
	fmt.Println(v)
}

//To access the field X of a struct when we have the struct pointer p we could write (*p).X.
//However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
