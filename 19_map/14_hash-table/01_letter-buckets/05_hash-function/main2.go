package main

import "fmt"

func main() {
	result := SecondFunc("Fck", 100)
	fmt.Println(result)
	// or
	fmt.Println(SecondFunc("Fck", 100))
}

func SecondFunc(word string, number int) int {
	a := int(word[0])
	b := number
	c := a + b
	return c
}

// a turns the word into int and [0] means that it will be the first letter.
// So for the word Fuck it's F.
// c means that the first letter of the word (in this case it's F)
// will add to the number (int) (in this case number is 100) and the result will be c
// 70 + 100 = 170
