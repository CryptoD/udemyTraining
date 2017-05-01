package main

import "fmt"

func name() float64 {
	var a float64 = 100
	return a + 10
}

func main() {

	var x float64
	fmt.Println("Number: ")
	fmt.Scan(&x)
	fmt.Println("Number two: ")

	var z = name()

	fmt.Println(z*100)
}