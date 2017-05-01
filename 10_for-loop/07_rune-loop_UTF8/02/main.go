package main

import "fmt"

func main() {
	for i := 500; i <= 520; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
