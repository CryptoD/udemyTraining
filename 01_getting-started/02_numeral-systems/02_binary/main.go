package main

import "fmt"

func main() {
	fmt.Printf("Binary: %b -- Hex: %x -- De: %d \n", 42, 42, 42)
	fmt.Printf("Binary: %b -- Hex: %X -- De: %d \n", 42, 42, 42)
	fmt.Printf("Binary: %#b -- Hex: %#x -- De: %#d \n", 42, 42, 42)
}
