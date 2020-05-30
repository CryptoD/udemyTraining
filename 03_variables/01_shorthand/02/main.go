package main

import "fmt"

func main() {

	a := 10
	b := "Text"
	c := 0.3
	d := true

	fmt.Printf("%v is a type %T.", a, a) //First I made it in 2 lines like below
	// Then I made it into 1 line.
	fmt.Printf("%v is a type %T.", b, b)

	fmt.Print(c)
	fmt.Printf(" is a type %T.", c)

	fmt.Print(d)
	fmt.Printf(" is a type %T.", d)
}
