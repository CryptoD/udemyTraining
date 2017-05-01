package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {

		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	for j := 0; j < 10; j++ {

		fmt.Printf("%d \t %b \t %x \t %q \n", j, j, j, j)
	}

}
