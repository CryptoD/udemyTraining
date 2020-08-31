package main

import "fmt"

func main() {
	var number1 int64
	var number2 int64

	fmt.Println("Chose one number and press ENTER, then chose another number and press ENTER")
	fmt.Scanln(&number1)
	fmt.Scanln(&number2)
	fmt.Println("Wybrane numery to", number1, "oraz", number2)
	fmt.Println(number1, "/", number2, "=", number1/number2)
	fmt.Println(number1, "*", number2, "=", number1*number2)
}
