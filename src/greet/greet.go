package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(sum())
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func sum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
