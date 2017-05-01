package main

import "fmt"

func main() {

	a := 43
	b := 1000
	c := "Hello"

	fmt.Println("normalne ", a, b, c)
	fmt.Println("memory address - ", &a, &b, &c)
	fmt.Printf("%d \n, %d \n, %d \n", &a, &b, &c)
}
