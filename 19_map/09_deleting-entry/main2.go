package main

import "fmt"

func main() {

	testujemy := make(map[int]string)

	testujemy[1] = "One"
	testujemy[2] = "Two"
	testujemy[3] = "Three"
	testujemy[4] = "Four"
	testujemy[3] = "Fixed Three"
	fmt.Println(testujemy)

	delete(testujemy, 2)
	fmt.Println(testujemy)

}
