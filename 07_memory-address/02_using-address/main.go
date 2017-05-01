package main

import "fmt"

var kmtom float64 = 1000

func main()  {

	var km float64
	fmt.Println("Wpisz kwote w km:")
	fmt.Scan(&km)
	var m = km * kmtom

	fmt.Println("Po przeliczeniu wychodzi ", m, "metrow.")
	fmt.Printf("Adres wyniku w pamieci RAM: %d", &m)

}