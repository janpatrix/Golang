package main

import (
	"fmt"
	"math"
)

func main() {
	a, b := Sqrt(2)
	fmt.Printf("The result is %f and it took %d iterations\n", a, b)
	fmt.Println("The real result is ", math.Sqrt(2))
}

func Sqrt(x float64) (float64, int) {

	z := x

	for i := 0; i < 10; i++ {
		y := z
		z -= (z*z - x) / (2 * z)
		if y == z {
			return y, i
		}
		fmt.Println(z)
	}
	return z, 10
}
