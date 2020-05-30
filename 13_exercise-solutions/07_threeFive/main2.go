package main

import "fmt"

func main() {
	counter := 0
	for a := 0; a < 20000; a++ {
		if a%4 == 0 {
			counter += a

		} else if a%5 == 0 {
			counter += a

		} else if a%6 == 0 {
			counter += a
		}

	}
	fmt.Println(counter)
}
