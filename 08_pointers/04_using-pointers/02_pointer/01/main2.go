package main

import "fmt"

func main() {

	x := 100
	y := &x

	fmt.Println("Literka X to:", x)
	fmt.Println("Literka Y to:", y)

	*y = 200

	fmt.Println("Nowa wartosc X to:", x)
	fmt.Println("Nowa wartosc Y to:", y)

}
