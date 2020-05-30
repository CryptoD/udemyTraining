package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
	Exported    int
}

func main() {
	p1 := person{"James", "Bond", 20, 007, 999}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	xyz := (string(bs))
	fmt.Println(xyz)
}

// I was testing different settings in this file. Finally found what was causing auto
// formatting and constantly changing location of where I type. It was UI settings (not
// Go Plus settings, those are fine.
