

import (
	"fmt"
	"math"
)

func swapword(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("============= Slide 1 : variable ====================")

	var i int
	var f float64
	var b bool
	var s string
	var r rune
	fmt.Printf("%v %v %v %v %q\n", i, f, b, r, s)

	fmt.Print("\n============= Slide 2 : convert type ====================\n")

	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y, z)

	fmt.Print("\n============= Slide 3 : function return multiple value ====================\n")

	s1, s2 := swapword("Toddy", "mon")
	fmt.Println(s1, s2)
}
