package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var delta = 1e-6
	z := x
	n := 0.0
	for math.Abs(n-z) > delta {
		n, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	const x = 2
	mine, theirs := Sqrt(x), math.Sqrt(x)
	fmt.Println(mine, theirs, mine-theirs)
}

//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//func Sqrt(x float64) float64 {
//	/* TODO */
//	prev := 0.0
//	z := x
//	for math.Abs(z-prev) > 1e-6 {
//		prev = z
//		z -= (z*z - x) / (2 * z)
//	}
//	return z
//}
//
//func main() {
//	fmt.Println(Sqrt(2))
//}
