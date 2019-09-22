package main

import "fmt"

func main() {
	fmt.Println("What's your name?")
	var a string
	var b string
	fmt.Scan(&a, &b)

	if a == "Bobby" {
		fmt.Println("Your name is", a)
	} else {
		fmt.Println("Mate, your name is Bobby!")
	}
}
