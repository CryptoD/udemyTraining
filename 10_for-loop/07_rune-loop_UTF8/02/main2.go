package main

import "fmt"

func main() {

	for a := 10000; a <= 10010; a++ {

		fmt.Printf("Normal %v, string %v, bit %v \n", a, string(a), []byte(string(a)))
	}
}
