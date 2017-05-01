package main

import (
	"fmt"
)

func main() {

	a := 100
	fmt.Println("A equals ", a)
	fmt.Println("A's address in RAM is ", &a)

	var b *int = &a

	fmt.Println("Result: ", b)
	fmt.Printf("Result in our language: %d", b)
}