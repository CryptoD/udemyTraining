package main

import "fmt"

func main() {


	for number := 1; number <= 100; number++ {


			if number % 3 == 0 && number % 5 == 0 {

				fmt.Println(number, "- FuzzBuzz")
			} else if number % 3 == 0 {

				fmt.Println(number, "- Buzz")
			} else if number % 5 == 0 {
				fmt.Println(number, "- Fuzz")
			}

	}

}
