package main

import (
	"fmt"
	"math"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	var r rune
	fmt.Printf("%v %v %v %v %q\n", i, f, b, r, s)

	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y, z)
}
