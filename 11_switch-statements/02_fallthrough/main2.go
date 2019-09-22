package main

import "fmt"

func main() {
	switch "Three" {
	case "One":
		fmt.Println("1")
	case "Two":
		fmt.Println("2")
	case "Three":
		fmt.Println("3")
    fallthrough
	case "Four":
		fmt.Println("4")
	case "Five":
		fmt.Println("5")
	}
}
