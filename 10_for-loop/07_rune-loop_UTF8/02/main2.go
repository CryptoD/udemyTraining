package main

import "fmt"
func main() {

  for a := 1000; a <= 1010; a++ {

  fmt.Printf("Normal %v, string %v, bit %v \n", a, string(a), []byte(string(a)))
}
}
