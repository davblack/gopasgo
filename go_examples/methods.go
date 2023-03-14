package main

import (
	"fmt"
	"math"
)

/**
** Go does not have classes. However, you can define methods on types.
*! A method is a function with a special receiver argument.
*  The receiver appears in its own argument list between the func keyword and the method name.
*  In this example, the Abs method has a receiver of type Vertex named v.
*? Remember: a method is just a function with a receiver argument.

*  You can declare a method on non-struct types, too.
*  You can only declare a method with a receiver whose type is defined in the same package as the method.
*! You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
 */

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
