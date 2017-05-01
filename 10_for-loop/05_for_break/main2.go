package main

import "fmt"


func main() {

	for a := 0; a <= 20; a++ {
		fmt.Printf("%d ", a)	// czyli "a" dodaje ciagle 1

		if a == 5 {
			fmt.Print("Five ")    // ale kiedy a wyniesie 5 jest Five
		}

		if a == 10 {		// jednak kiedy a rowne 19 to przerwij
			break
		}
	}
}