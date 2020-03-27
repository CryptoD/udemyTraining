package main

import "fmt"

func main() {
	first()
	second()
}

func first() {
	for b := 1; b < 10; b++ {
		fmt.Println("Option A:", b)
	}
}

func second() {
	for c := 1; c < 10; c++ {
		fmt.Println("Option B: ", c)
	}
}