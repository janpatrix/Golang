package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var c, python, java bool = true, false, true
var i int = 15

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("The number Pi is: ", math.Pi)
	fmt.Println("5 + 4 = ", add(5, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(i, c, python, java)
}

func add(x, y int) int {
	return (x + y)
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
