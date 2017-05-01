package main

import "fmt"

func main() {

	var a = 100
	var b *int = &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
}