package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{1, 114, 8, 121, 430, 55, 66}
	fmt.Println("Random numbers:", s)
	sort.Ints(s)
	fmt.Println("After a bit of coding:", s)
}
