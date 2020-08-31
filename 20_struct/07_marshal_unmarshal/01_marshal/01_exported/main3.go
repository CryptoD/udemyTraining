package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	parameter1 int
	Parameter2 int
	parameter3 string
	parameter4 bool
}

func main() {
	p1 := person{77, 88, "Hello", true}
	p2, _ := json.Marshal(p1)
	// p3 := string(p1) <-- impossible to change the whole of p1 into string
	p4 := string(p2)
	fmt.Println(p1)
	fmt.Println(p2)
	// fmt.Println(p3)
	fmt.Println(p4)
}
