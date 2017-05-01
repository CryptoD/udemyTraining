package main

import "fmt"

func main() {

	var name string

	fmt.Println("What's your name fella?")
	fmt.Scan(&name)
	fmt.Println("How you doing", name, "?")
}