package main

import "fmt"

func main() {

	var one int
	var two int

	fmt.Print("Enter your number: ")
	fmt.Scan(&one)
	fmt.Print("And another number: ")
	fmt.Scan(&two)

	fmt.Println("Thanks for your numbers!", one, "devided by", two, "is", one / two)
}