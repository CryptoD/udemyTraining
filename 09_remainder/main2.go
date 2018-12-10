package main

import (
	"fmt"
)

func main() {

	a := 10
	b := 3
	c := a / b

	fmt.Println(c) // Shows the result but it's simplified. There's some reminder to it

	d := a % b // This is a reminder

	fmt.Println(d)

	// Now let's write something for a human to understand

	fmt.Println("Number", a, " devided by number", b, "is", c, "and the reminder is", d)
}
