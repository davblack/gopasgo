package main

import "fmt"

//Map literals are like struct literals, but the keys are required.
//If the top-level type is just a type name, you can omit it from the elements of the literal.
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
