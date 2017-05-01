package main

import "fmt"

func main() {

	for x := 0; x <= 60; x++ {

		if x % 15 == 0 {
			fmt.Println(x, "-- BuzzFizz")
		} else if x % 5 == 0 {
			fmt.Println(x, "-- BUZZ")
		} else if x % 3 == 0 {
			fmt.Println(x, "-- FIZZ")
		} else {
			fmt.Println(x)
		}
	}
}