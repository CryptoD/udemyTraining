package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("Put your number: ")
	fmt.Scan(&num1)
	fmt.Print("Put number 2: ")
	fmt.Scan(&num2)
	num3 := num1 / num2
	fmt.Println("Your new number is: ", num3)
}