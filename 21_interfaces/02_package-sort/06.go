package main
import "fmt"
import "sort"

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}