// This example shows how to unmarshall more than one thing

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   float64
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := Person{
		First: "Emma",
		Last:  "Moneypenn",
		Age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	xp := []Person{p1, p2}

	fmt.Println(xp)

	bs, e := json.Marshal(xp)
	if e != nil {
		panic(e)
	}

	fmt.Println(bs)

	var xp2 []Person

	json.Unmarshal(bs, &xp2)

	fmt.Println(xp)
}
