package main

import (
	"fmt"
)

func first(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("Numbers")

	case string:
		fmt.Println("Words")
	}
}

	func main() {

		first("ABC")
		first(7)

	}