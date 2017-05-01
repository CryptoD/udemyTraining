package main

import "fmt"

func main() {

	var v = calculate(10, 20, 30, 40, 50)
	fmt.Println(v)
}

func calculate(x ...float64) float64 {

	fmt.Println(x)			// Pokazujemy liczby
	var total float64		// Total bedzie sumowac ilosc cyfr
	for _, z := range x {		//
		total += z
	}
	return total / float64(len(x))
}