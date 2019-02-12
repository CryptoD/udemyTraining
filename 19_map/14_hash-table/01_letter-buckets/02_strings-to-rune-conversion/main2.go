package main

import "fmt"

func main() {
	letter := rune("ABA"[2])
	fmt.Println(letter)
}

// It's done in order.
// Number 0 is representing A, nuber 1 represents B, number 3 represents A
// So becuase it's 3, it represents string B, transforms into ASCII (66)
