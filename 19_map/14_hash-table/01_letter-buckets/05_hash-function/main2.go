package main

import "fmt"

func main() {
	result := SecondFunc("Fuck", 100)
	fmt.Println(result)
	// or
	fmt.Println(SecondFunc("Fuck", 100))
}

func SecondFunc(word string, number int) int {
	a := int(word[0])
	b := number
	c := a + b
	return c
}
