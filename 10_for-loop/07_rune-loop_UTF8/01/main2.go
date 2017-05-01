package main

import "fmt"

func main() {

for x := 1000; x < 1003; x++ {

	fmt.Println(x, " - ", string(x), " - ", []byte(string(x)))

	}

h := "中"

	fmt.Println(h, " - ", string(h), " - ", []byte(string(h)))


}