package main

import (
	"fmt"
	"math"
)

type Vertext struct {
	X, Y float64
}

func (v Vertext) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertext) Addition() float64 {
	return (v.X + v.Y)
}

func (v Vertext) printHello() {
	fmt.Println("Hello!")

}

func main() {
	v := Vertext{3, 4}
	fmt.Println("The ABS result is ", v.Abs())
	fmt.Println("The SUM result is ", v.Addition())
	v.printHello()
}
