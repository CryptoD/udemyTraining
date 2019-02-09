package main

import "fmt"

func main() {

	m := make(map[int]int)

	m[1] = 7
	m[2] = 2
	fmt.Println(m)
}

// Prints something like this:
// map[1:7 2:2]
