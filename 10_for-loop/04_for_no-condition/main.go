package main

import "fmt"

func main() {
	i := 0
	for {
		a := &i
		fmt.Printf("\r %d - %x - %x", i, a)
		i++
	}
}
