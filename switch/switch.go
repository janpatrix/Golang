package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	switchOS()
	isItSaturday()
	goodMorning()

}

func switchOS() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func isItSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 0:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

func goodMorning() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	case t.Hour() < 22:
		fmt.Println("Good evening")
	default:
		fmt.Println("Good night")
	}
}
