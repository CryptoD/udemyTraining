package main
import (
	"fmt"
	"sync"
)

var anyname sync.WaitGroup

func main() {
	anyname.Add(2)
	go first()
	go second()
	anyname.Wait()
}

func first() {
	for a := 1; a < 10; a++ {
		fmt.Println("Option One: ", a)
	}
	anyname.Done()
}

func second() {
	for b := 1; b < 10; b++ {
		fmt.Println("Option Two: ", b)
	}
	anyname.Done()
}