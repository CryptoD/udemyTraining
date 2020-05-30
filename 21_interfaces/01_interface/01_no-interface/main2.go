package main

import "fmt"

type lake struct {
	area int
}

func (depth lake) total() int {
	return depth.area * depth.area
}
func main() {
	x := lake{10}
	fmt.Println("Total lake power is: ", x.total())
}
