package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	var s []int
	printMoreSlice(s)

	s = append(s, 0)
	printMoreSlice(s)

	s = append(s, 1)
	printMoreSlice(s)

	s = append(s, 2, 3, 4)
	printMoreSlice(s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap =%d %v\n", s, len(x), cap(x), x)
}

func printMoreSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func powPow() {

	var pow = []int{1, 2, 4, 8, 18, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
