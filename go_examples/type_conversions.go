package main

import (
	"fmt"
	"math"
)

//Unlike in C, in Go assignment between items of different type requires an explicit conversion.
//Try removing the float64 or uint conversions in the example and see what happens.
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
