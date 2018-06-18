package main

import "fmt"
func main() {
 a := []int{1, 3, 5, 500}
 b = a + 3
 fmt.Print(a, b)

}

// Wrong because a is []int and b is normal int. Shows error
