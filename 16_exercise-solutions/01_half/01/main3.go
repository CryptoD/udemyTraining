package main

import "fmt"

func calculation(number int) (int, bool) {
	return number / 2, number%2 == 0

}
func main() {
	fmt.Println(calculation(number(5)))
}

// this one is wrong because you don't put shit into "number"
// "Number" is being created so you just need calculation(put shit here)
