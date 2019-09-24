package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(dd ...float64) float64 {
	fmt.Println(dd)
	fmt.Println(len(dd))
	fmt.Printf("%T \n", dd)
	var total float64
	for _, v := range dd {
		total += v
	}
	return total / float64(len(dd))

}
