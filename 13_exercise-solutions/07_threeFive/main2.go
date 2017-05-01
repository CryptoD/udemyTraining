package main

import "fmt"

func main() {

	x := 0

	for y := 0; y < 1000; y++ {
		if y%3 == 0 {
			x += y
		} else if y%5 == 0 {
			x += y
		}
	}
	fmt.Print(x)
}