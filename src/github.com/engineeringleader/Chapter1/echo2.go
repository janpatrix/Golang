package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//s, sep := "", " "
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])
}
