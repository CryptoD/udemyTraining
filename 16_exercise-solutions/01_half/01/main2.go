package main

import "fmt"

func one(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(one(5))
}
// func one creates "n" that is expressed in two numbers: int and bool.
// Then in return you just put whatever you want them to be
