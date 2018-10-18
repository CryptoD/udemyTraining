package main

import "fmt"

func count(numbers ...int) int {
	var number1 int
	for _, number2 := range numbers {
		if number2 > number1 {
			number1 = number2
		}
	}
	return number1
}
func main() {

	number3 := count(1, 2, 5, 50)
	fmt.Println(number3)
}
